package public

import (
	"context"
	errors2 "errors"

	"github.com/gstones/moke-kit/utility"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	leaderboard "github.com/moke-game/platform/api/gen/leaderboard/api"
	"github.com/moke-game/platform/services/leaderboard/internal/service/errors"
)

func (s *Service) GetLeaderboard(
	ctx context.Context,
	request *leaderboard.GetLeaderboardRequest,
) (*leaderboard.GetLeaderboardResponse, error) {
	page := request.GetPage()
	pageSize := request.GetPageSize()
	if page < 1 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	start := (page - 1) * pageSize
	end := page*pageSize - 1
	if scores, err := s.db.GetTopWithScores(request.GetId(), int64(start), int64(end)); err != nil {
		return nil, err
	} else if len(scores) == 0 {
		return &leaderboard.GetLeaderboardResponse{}, nil
	} else {
		uids := make([]string, 0)
		for _, v := range scores {
			uids = append(uids, v.Member.(string))
		}
		stars := s.db.GetLeaderboardStars(ctx, request.GetId(), uids...)
		entries := s.makeLeaderboardEntries(scores, stars)
		return &leaderboard.GetLeaderboardResponse{
			Entries: entries,
		}, nil
	}
}

func (s *Service) makeLeaderboardEntries(datas []redis.Z, stars map[string]int64) []*leaderboard.LeaderboardEntry {
	entries := make([]*leaderboard.LeaderboardEntry, len(datas))
	for k, v := range datas {
		member := v.Member.(string)
		star := stars[member]
		entries[k] = &leaderboard.LeaderboardEntry{
			Uid:   member,
			Score: v.Score,
			Star:  star,
		}
	}
	return entries
}

func (s *Service) GetRank(ctx context.Context, request *leaderboard.GetRankRequest) (*leaderboard.GetRankResponse, error) {
	member := ""
	if uid, ok := utility.FromContext(ctx, utility.UIDContextKey); !ok {
		return nil, errors.ErrNoMetaData
	} else {
		member = uid
	}
	if request.GetCountry() != "" {
		member = request.GetCountry()
	}

	if rank, err := s.db.GetRankScore(request.GetId(), member); err != nil {
		if errors2.Is(err, redis.Nil) {
			return &leaderboard.GetRankResponse{}, nil
		}
		s.logger.Error("get rank failed", zap.Error(err), zap.String("member", member), zap.String("id", request.GetId()))
		return nil, errors.ErrGeneralFailure
	} else {
		return &leaderboard.GetRankResponse{
			Rank:  rank.Rank + 1,
			Score: rank.Score,
		}, nil
	}
}

func (s *Service) StarLeaderboard(ctx context.Context, request *leaderboard.StarLeaderboardRequest) (*leaderboard.StarLeaderboardResponse, error) {
	num, err := s.db.StarLeaderboard(ctx, request.GetId(), request.GetUid())
	if err != nil {
		s.logger.Error("star leaderboard failed", zap.Error(err), zap.String("uid", request.GetUid()), zap.String("id", request.GetId()))
		return nil, errors.ErrGeneralFailure
	}
	return &leaderboard.StarLeaderboardResponse{
		StarCount: num,
	}, nil
}
