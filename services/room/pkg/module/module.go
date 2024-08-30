package module

import (
	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/room/internal"
	"github.com/moke-game/platform/services/room/internal/common"
	"github.com/moke-game/platform/services/room/pkg/rfx"
)

// RoomModule  backend for frontend service module
var RoomModule = fx.Module("room", fx.Options(
	rfx.SettingsModule,
	internal.Module,
	globalModule,
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("room")
	}),
))

var globalModule = fx.Invoke(
	func(
		//l *zap.Logger,
		//p cfx.ConfigsParams,
		aParams mfx.AppParams,
	) {
		common.DeploymentGlobal = utility.ParseDeployments(aParams.Deployment)
	},
)
