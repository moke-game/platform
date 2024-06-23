package public

import (
	"time"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/utility"

	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/pkg/ofx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"

	"github.com/moke-game/platform/services/auth/pkg/afx"

	pb "github.com/moke-game/platform/api/gen/auth"
	"github.com/moke-game/platform/services/auth/service/db"
)

type Service struct {
	utility.WithoutAuth
	logger    *zap.Logger
	jwtSecret string
	url       string
	db        *db.Database
	tracer    trace.Tracer
	redisCli  *redis.Client
	jwtExpire time.Duration
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterAuthServiceServer(server.GrpcServer(), s)
	return nil
}

func NewService(
	l *zap.Logger,
	jwtSecret string,
	url string,
	coll diface.ICollection,
	redisCache diface.ICache,
	appName string,
	tracer trace.Tracer,
	redisCli *redis.Client,
	jwtExpire int32,
) (result *Service, err error) {
	result = &Service{
		logger:    l,
		jwtSecret: jwtSecret,
		url:       url,
		db:        db.OpenDatabase(l, appName, coll, redisCache),
		tracer:    tracer,
		redisCli:  redisCli,
		jwtExpire: time.Duration(jwtExpire) * time.Hour,
	}
	return
}

var ServiceModule = fx.Provide(
	func(
		l *zap.Logger,
		sSetting afx.AuthSettingParams,
		aParams mfx.AppParams,
		dbProvider ofx.DocumentStoreParams,
		rcParams ofx.RedisCacheParams,
		redisParams ofx.RedisParams,
	) (rpc sfx.GrpcServiceResult, err error) {
		if coll, e := dbProvider.DriverProvider.OpenDbDriver(sSetting.AuthStoreName); e != nil {
			err = e
		} else {
			if svc, e := NewService(
				l,
				sSetting.JwtTokenSecret,
				sSetting.AuthUrl,
				coll,
				rcParams.RedisCache,
				aParams.AppName,
				otel.GetTracerProvider().Tracer(sSetting.AuthStoreName),
				redisParams.Redis,
				sSetting.JwtTokenExpire,
			); e != nil {
				err = e
			} else {
				rpc.GrpcService = svc
			}
		}
		return
	},
)
