package afx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"

	pb "github.com/gstones/platform/api/gen/auth"
)

type AuthClientParams struct {
	fx.In

	AuthClient pb.AuthServiceClient `name:"AuthClient"`
}

type AuthClientResult struct {
	fx.Out

	AuthClient pb.AuthServiceClient `name:"AuthClient"`
}

func NewAuthClient(host string, logger *zap.Logger, sSetting sfx.SecuritySettingsParams) (pb.AuthServiceClient, error) {
	if sSetting.MTLSEnable {
		if conn, err := tools.DialWithSecurity(
			host,
			sSetting.ClientCert,
			sSetting.ClientKey,
			sSetting.ServerName,
			sSetting.ServerCaCert,
		); err != nil {
			return nil, err
		} else {
			return pb.NewAuthServiceClient(conn), nil
		}
	} else {
		if conn, err := tools.DialInsecure(host); err != nil {
			return nil, err
		} else {
			return pb.NewAuthServiceClient(conn), nil
		}
	}
}

var AuthClientModule = fx.Provide(
	func(
		setting AuthSettingParams,
		sSetting sfx.SecuritySettingsParams,
		logger *zap.Logger,
	) (out AuthClientResult, err error) {
		if cli, e := NewAuthClient(setting.AuthUrl, logger, sSetting); e != nil {
			err = e
		} else {
			out.AuthClient = cli
		}
		return
	},
)
