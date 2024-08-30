package private

import (
	"context"

	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"

	pb "github.com/moke-game/platform/api/gen/knapsack/api"
	"github.com/moke-game/platform/services/knapsack/changes"
	"github.com/moke-game/platform/services/knapsack/errors"
)

func (s *Service) AddItem(_ context.Context, request *pb.AddItemPrivateRequest) (*pb.AddItemPrivateResponse, error) {
	if dao, err := s.db.LoadOrCreateKnapsack(request.Uid); err != nil {
		s.logger.Error("load knapsack failed", zap.Error(err))
		return nil, errors.ErrNotFound
	} else if err := dao.Update(func() bool {
		if len(request.GetFeatures()) > 0 {
			dao.AddFeatures(request.GetFeatures())
		}
		if len(request.GetItems()) > 0 {
			dao.AddItems(request.GetItems())
		}
		return true
	}); err != nil {
		s.logger.Error("add item failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	} else if err := s.pubChanges(request.Uid, dao.GetAndDeleteChanges(request.Items, nil, request.Source)); err != nil {
		s.logger.Error("publish knapsack changes failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.AddItemPrivateResponse{}, nil
}

func (s *Service) pubChanges(uid string, modify *pb.KnapsackModify) error {
	if modify.Knapsack == nil {
		return nil
	}
	topic := changes.CreateTopic(uid)
	if data, err := changes.Pack(modify); err != nil {
		s.logger.Error("marshal knapsack failed", zap.Error(err))
		return err
	} else if err := s.mq.Publish(topic, miface.WithBytes(data)); err != nil {
		s.logger.Error("publish knapsack changes failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) RemoveItem(_ context.Context, request *pb.RemoveItemPrivateRequest) (*pb.RemoveItemPrivateResponse, error) {
	if dao, err := s.db.LoadKnapsack(request.Uid); err != nil {
		s.logger.Error("load knapsack failed", zap.Error(err))
		return nil, errors.ErrNotFound
	} else if err := dao.Update(func() bool {
		if err := dao.RemoveItems(request.GetItems()); err != nil {
			s.logger.Error("remove item failed", zap.Error(err))
			return false
		}
		return true
	}); err != nil {
		return nil, errors.ErrNotEnough
	} else if err := s.pubChanges(request.Uid, dao.GetAndDeleteChanges(nil, request.Items, request.Source)); err != nil {
		s.logger.Error("publish knapsack changes failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.RemoveItemPrivateResponse{}, nil
}

func (s *Service) GetItemById(_ context.Context, request *pb.GetItemByIdPrivateRequest) (*pb.GetItemByIdPrivateResponse, error) {
	if dao, err := s.db.LoadKnapsack(request.Uid); err != nil {
		s.logger.Error("load knapsack failed", zap.Error(err))
		return nil, errors.ErrNotFound
	} else {
		item := dao.Data.Items[request.ItemId]
		return &pb.GetItemByIdPrivateResponse{Item: item}, nil
	}

}

func (s *Service) GetKnapsack(_ context.Context, request *pb.GetKnapsackRequest) (*pb.GetKnapsackResponse, error) {
	if dao, err := s.db.LoadKnapsack(request.Uid); err != nil {
		s.logger.Error("load knapsack failed", zap.Error(err))
		return nil, errors.ErrNotFound
	} else {
		return &pb.GetKnapsackResponse{Knapsack: dao.Data}, nil
	}

}
