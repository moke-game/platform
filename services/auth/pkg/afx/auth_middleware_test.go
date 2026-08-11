package afx

import (
	"context"
	"errors"
	"testing"

	"github.com/gstones/moke-kit/utility"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/moke-game/platform/api/gen/auth/api"
)

type mockAuthClient struct {
	uid string
	err error
}

func (m *mockAuthClient) Authenticate(context.Context, *pb.AuthenticateRequest, ...grpc.CallOption) (*pb.AuthenticateResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Authenticate")
}
func (m *mockAuthClient) RefreshToken(context.Context, *pb.RefreshTokenRequest, ...grpc.CallOption) (*pb.RefreshTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "RefreshToken")
}
func (m *mockAuthClient) ValidateToken(_ context.Context, in *pb.ValidateTokenRequest, _ ...grpc.CallOption) (*pb.ValidateTokenResponse, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &pb.ValidateTokenResponse{Uid: m.uid}, nil
}
func (m *mockAuthClient) PackToken(context.Context, *pb.PackTokenRequest, ...grpc.CallOption) (*pb.PackTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "PackToken")
}
func (m *mockAuthClient) ClearToken(context.Context, *pb.ClearTokenRequest, ...grpc.CallOption) (*pb.ClearTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ClearToken")
}
func (m *mockAuthClient) Delete(context.Context, *pb.DeleteRequest, ...grpc.CallOption) (*pb.DeleteResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Delete")
}
func (m *mockAuthClient) AddBlocked(context.Context, *pb.AddBlockedRequest, ...grpc.CallOption) (*pb.AddBlockedResponse, error) {
	return nil, status.Error(codes.Unimplemented, "AddBlocked")
}

type mockStream struct {
	method string
}

func (m *mockStream) Method() string                  { return m.method }
func (m *mockStream) SetHeader(metadata.MD) error     { return nil }
func (m *mockStream) SendHeader(metadata.MD) error    { return nil }
func (m *mockStream) SetTrailer(metadata.MD) error    { return nil }

func ctxWithMethodAndAuth(method, bearer string) context.Context {
	ctx := grpc.NewContextWithServerTransportStream(context.Background(), &mockStream{method: method})
	if bearer != "" {
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("authorization", "bearer "+bearer))
	}
	return ctx
}

func TestAuthorAuth(t *testing.T) {
	t.Parallel()

	const method = "/game0.pb.Game0Service/Hi"
	author := &Author{
		client:        &mockAuthClient{uid: "uid-1"},
		unAuthMethods: map[string]struct{}{},
	}

	t.Run("missing bearer fails", func(t *testing.T) {
		t.Parallel()
		_, err := author.Auth(ctxWithMethodAndAuth(method, ""))
		if err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("validate failure fails", func(t *testing.T) {
		t.Parallel()
		a := &Author{
			client:        &mockAuthClient{err: errors.New("bad token")},
			unAuthMethods: map[string]struct{}{},
		}
		_, err := a.Auth(ctxWithMethodAndAuth(method, "tok"))
		if err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("success sets uid", func(t *testing.T) {
		t.Parallel()
		ctx, err := author.Auth(ctxWithMethodAndAuth(method, "tok"))
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		uid, ok := ctx.Value(utility.UIDContextKey).(string)
		if !ok || uid != "uid-1" {
			t.Fatalf("uid=%q ok=%v", uid, ok)
		}
	})

	t.Run("unauth method bypasses validate", func(t *testing.T) {
		t.Parallel()
		a := &Author{
			client:        &mockAuthClient{err: errors.New("should not call")},
			unAuthMethods: map[string]struct{}{},
		}
		a.AddUnAuthMethod(method)
		ctx, err := a.Auth(ctxWithMethodAndAuth(method, ""))
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if v, _ := ctx.Value(utility.WithOutTag).(bool); !v {
			t.Fatal("expected WithOutTag")
		}
	})
}
