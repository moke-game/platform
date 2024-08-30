package private

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	leaderboard "github.com/moke-game/platform/api/gen/leaderboard/api"
	"github.com/moke-game/platform/services/leaderboard/internal/service/errors"
)

func (s *Service) ClearLeaderboard(ctx context.Context, request *leaderboard.ClearLeaderboardRequest) (*leaderboard.ClearLeaderboardResponse, error) {
	if err := s.db.ClearLeaderboard(request.GetId()); err != nil {
		s.logger.Error("clear leaderboard error", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	return &leaderboard.ClearLeaderboardResponse{}, nil
}

func (s *Service) ExpireLeaderboard(
	_ context.Context,
	request *leaderboard.ExpireLeaderboardRequest,
) (*leaderboard.ExpireLeaderboardResponse, error) {
	expire := s.expireTime
	if request.ExpireTime != 0 {
		expire = time.Duration(request.ExpireTime) * time.Hour * 24
	}
	isDelete := s.db.ExpiredLeaderboard(request.GetId(), expire)
	if !isDelete {
		return &leaderboard.ExpireLeaderboardResponse{IsDeleted: isDelete}, nil
	}
	scores, err := s.db.GetTopWithScores(request.GetId(), int64(0), int64(request.Num-1))
	if err != nil {
		return nil, err
	}
	entries := s.makeLeaderboardEntries(scores)
	s.db.ExpireLeaderboardStar(request.GetId(), expire)
	return &leaderboard.ExpireLeaderboardResponse{IsDeleted: isDelete, Entries: entries}, nil
}

func (s *Service) makeLeaderboardEntries(datas []redis.Z) []*leaderboard.LeaderboardEntry {
	entries := make([]*leaderboard.LeaderboardEntry, len(datas))
	for k, v := range datas {
		entries[k] = &leaderboard.LeaderboardEntry{
			Uid:   v.Member.(string),
			Score: v.Score,
		}
	}
	return entries
}

func (s *Service) UpdateScore(ctx context.Context, request *leaderboard.UpdateScoreRequest) (*leaderboard.UpdateScoreResponse, error) {
	scores := request.Scores
	members := make([]string, 0, len(scores))
	for k := range scores {
		members = append(members, k)
	}
	before, err := s.db.MGetScores(request.GetId(), members)
	if err != nil {
		s.logger.Error("get scores error", zap.Error(err))
		return nil, errors.ErrGeneralFailure
	}
	oldTopOne := s.getTopOne(request.GetId())
	update := make(map[string]float64)
	if request.UpdateType == leaderboard.UpdateScoreRequest_ADD {
		for k, v := range before {
			if _, ok := scores[k]; ok {
				update[k] = v + scores[k]
			}
		}
	} else if request.UpdateType == leaderboard.UpdateScoreRequest_GTR {
		for k, v := range before {
			if _, ok := scores[k]; ok {
				if v < scores[k] {
					update[k] = scores[k]
				}
			}
		}
	} else if request.UpdateType == leaderboard.UpdateScoreRequest_LSR {
		for k, v := range before {
			if _, ok := scores[k]; ok {
				if v > scores[k] {
					update[k] = scores[k]
				}
			}
		}
	}

	scoresTimestamp := make(map[string]float64)
	for k, v := range update {
		sc := s.addTimestamp(v)
		scoresTimestamp[k] = sc
	}
	if err := s.db.UpdateScore(request.GetId(), scoresTimestamp); err != nil {
		return nil, err
	}
	curTopOne := s.getTopOne(request.GetId())
	res := &leaderboard.UpdateScoreResponse{}
	if len(oldTopOne) > 0 && oldTopOne != curTopOne {
		res.Id = request.GetId()
		res.OldUid = oldTopOne
		res.CurrentUid = curTopOne
	}
	return res, nil
}

var BaseTime = time.Date(2124, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()

// addTimestamp add timestamp to score decimal
func (s *Service) addTimestamp(score float64) float64 {
	decimal := float64(BaseTime-time.Now().UnixMilli()) / float64(BaseTime)
	return score + decimal
}

func (s *Service) getTopOne(id string) string {
	//查出原本的第一名
	uid := ""
	amount, er := s.db.GetLeaderboardAmount(id)
	if er != nil {
		s.logger.Error("GetLeaderboardAmount error", zap.Error(er))
	} else if amount > 0 {
		withScores, err := s.db.GetTopWithScores(id, 0, 0)
		if err != nil {
			s.logger.Error("GetTopWithScores error", zap.Error(err))
		} else if len(withScores) > 0 {
			uid = withScores[0].Member.(string)
		}
	}
	return uid
}
