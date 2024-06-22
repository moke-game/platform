package public

import (
	"context"

	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/zap"

	pb "github.com/gstones/platform/api/gen/knapsack"
	"github.com/gstones/platform/services/knapsack/changes"
	"github.com/gstones/platform/services/knapsack/errors"
)

func (s *Service) GetKnapsack(ctx context.Context, _ *pb.GetKnapsackRequest) (*pb.GetKnapsackResponse, error) {
	if uid, ok := utility.FromContext(ctx, utility.UIDContextKey); !ok {
		return nil, errors.ErrNoMetaData
	} else if dao, err := s.db.LoadKnapsack(uid); err != nil {
		return nil, err
	} else {
		return &pb.GetKnapsackResponse{
			Knapsack: dao.ToProto(),
		}, nil
	}
}

func (s *Service) AddItem(ctx context.Context, request *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	if uid, ok := utility.FromContext(ctx, utility.UIDContextKey); !ok {
		s.logger.Error("get uid from context err")
		return nil, errors.ErrNoMetaData
	} else if dao, err := s.db.LoadOrCreateKnapsack(uid); err != nil {
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
	} else if err := s.pubChanges(uid, dao.GetAndDeleteChanges(request.Items, nil, request.Source)); err != nil {
		s.logger.Error("publish knapsack changes failed", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &pb.AddItemResponse{}, nil
}

func (s *Service) pubChanges(uid string, knapsack *pb.KnapsackModify) error {
	if knapsack.Knapsack == nil {
		return nil
	}
	topic := changes.CreateTopic(uid)
	if data, err := changes.Pack(knapsack); err != nil {
		s.logger.Error("marshal knapsack failed", zap.Error(err))
		return err
	} else if err := s.mq.Publish(topic, miface.WithBytes(data)); err != nil {
		s.logger.Error("publish knapsack changes failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) RemoveItem(ctx context.Context, request *pb.RemoveItemRequest) (*pb.RemoveItemResponse, error) {
	if uid, ok := utility.FromContext(ctx, utility.UIDContextKey); !ok {
		return nil, errors.ErrNoMetaData
	} else if dao, err := s.db.LoadKnapsack(uid); err != nil {
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
	} else {
		if err := s.pubChanges(uid, dao.GetAndDeleteChanges(nil, request.Items, request.Source)); err != nil {
			s.logger.Error("publish knapsack changes failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
	}
	return &pb.RemoveItemResponse{}, nil
}

func (s *Service) RemoveThenAddItem(ctx context.Context, request *pb.RemoveThenAddItemRequest) (*pb.RemoveThenAddItemResponse, error) {
	if uid, ok := utility.FromContext(ctx, utility.UIDContextKey); !ok {
		return nil, errors.ErrNoMetaData
	} else if dao, err := s.db.LoadKnapsack(uid); err != nil {
		s.logger.Error("load knapsack failed", zap.Error(err))
		return nil, errors.ErrNotFound
	} else if err := dao.Update(func() bool {
		if err := dao.RemoveItems(request.GetRemoveItems()); err != nil {
			s.logger.Error("remove item failed", zap.Error(err))
			return false
		}
		dao.AddItems(request.GetAddItems())
		return true
	}); err != nil {
		return nil, errors.ErrNotEnough
	} else {
		if err := s.pubChanges(uid, dao.GetAndDeleteChanges(request.AddItems, request.RemoveItems, request.Source)); err != nil {
			s.logger.Error("publish knapsack changes failed", zap.Error(err))
			return nil, errors.ErrGeneralFailure
		}
	}
	return &pb.RemoveThenAddItemResponse{}, nil
}
