package afx

import (
	"context"
	"testing"

	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
)

type publicSvc struct{}

func (publicSvc) RegisterWithGrpcServer(siface.IGrpcServer) error { return nil }

type privateSvc struct {
	utilityWithoutAuth
}

type utilityWithoutAuth struct{}

func (utilityWithoutAuth) AuthFuncOverride(ctx context.Context, _ string) (context.Context, error) {
	return ctx, nil
}

func (privateSvc) RegisterWithGrpcServer(siface.IGrpcServer) error { return nil }

type gatewaySvc struct{}

func (gatewaySvc) RegisterWithGatewayServer(siface.IGatewayServer) error { return nil }

type stubAuth struct{}

func (stubAuth) Auth(ctx context.Context) (context.Context, error) { return ctx, nil }

func (stubAuth) AddUnAuthMethod(string) {}

func TestCheckProdPublicAuthFailClosed(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		deploy  string
		auth    sfx.AuthMiddlewareParams
		grpc    sfx.GrpcServiceParams
		gateway sfx.GatewayServiceParams
		wantErr bool
	}{
		{
			name:   "non-prod allows missing middleware",
			deploy: "local",
			grpc:   sfx.GrpcServiceParams{GrpcServices: []siface.IGrpcService{publicSvc{}}},
		},
		{
			name:   "prod private-only ok without middleware",
			deploy: "prod",
			grpc:   sfx.GrpcServiceParams{GrpcServices: []siface.IGrpcService{privateSvc{}}},
		},
		{
			name:    "prod public without middleware fails",
			deploy:  "prod",
			grpc:    sfx.GrpcServiceParams{GrpcServices: []siface.IGrpcService{publicSvc{}}},
			wantErr: true,
		},
		{
			name:   "prod public with middleware ok",
			deploy: "prod",
			auth:   sfx.AuthMiddlewareParams{AuthMiddleware: stubAuth{}},
			grpc:   sfx.GrpcServiceParams{GrpcServices: []siface.IGrpcService{publicSvc{}}},
		},
		{
			name:   "prod mixed services require middleware",
			deploy: "prod",
			grpc: sfx.GrpcServiceParams{GrpcServices: []siface.IGrpcService{
				privateSvc{},
				publicSvc{},
			}},
			wantErr: true,
		},
		{
			name:    "prod gateway without middleware fails",
			deploy:  "prod",
			gateway: sfx.GatewayServiceParams{GatewayServices: []siface.IGatewayService{gatewaySvc{}}},
			wantErr: true,
		},
		{
			name:    "prod gateway with middleware ok",
			deploy:  "prod",
			auth:    sfx.AuthMiddlewareParams{AuthMiddleware: stubAuth{}},
			gateway: sfx.GatewayServiceParams{GatewayServices: []siface.IGatewayService{gatewaySvc{}}},
		},
		{
			name:   "prod pass-through with private-only ok",
			deploy: "prod",
			auth:   sfx.AuthMiddlewareParams{AuthMiddleware: &passThroughAuthor{}},
			grpc:   sfx.GrpcServiceParams{GrpcServices: []siface.IGrpcService{privateSvc{}}},
		},
		{
			name:    "prod pass-through with public fails",
			deploy:  "prod",
			auth:    sfx.AuthMiddlewareParams{AuthMiddleware: &passThroughAuthor{}},
			grpc:    sfx.GrpcServiceParams{GrpcServices: []siface.IGrpcService{publicSvc{}}},
			wantErr: true,
		},
		{
			name:    "prod pass-through with gateway fails",
			deploy:  "prod",
			auth:    sfx.AuthMiddlewareParams{AuthMiddleware: &passThroughAuthor{}},
			gateway: sfx.GatewayServiceParams{GatewayServices: []siface.IGatewayService{gatewaySvc{}}},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := CheckProdPublicAuthFailClosed(
				mfx.AppParams{Deployment: tc.deploy},
				tc.auth,
				tc.grpc,
				tc.gateway,
			)
			if tc.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
