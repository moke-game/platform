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
	uid           string
	err           error
	lastAccessTok string
}

func (m *mockAuthClient) Authenticate(context.Context, *pb.AuthenticateRequest, ...grpc.CallOption) (*pb.AuthenticateResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Authenticate")
}
func (m *mockAuthClient) RefreshToken(context.Context, *pb.RefreshTokenRequest, ...grpc.CallOption) (*pb.RefreshTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "RefreshToken")
}
func (m *mockAuthClient) ValidateToken(_ context.Context, in *pb.ValidateTokenRequest, _ ...grpc.CallOption) (*pb.ValidateTokenResponse, error) {
	m.lastAccessTok = in.GetAccessToken()
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

func (m *mockStream) Method() string               { return m.method }
func (m *mockStream) SetHeader(metadata.MD) error  { return nil }
func (m *mockStream) SendHeader(metadata.MD) error { return nil }
func (m *mockStream) SetTrailer(metadata.MD) error { return nil }

func ctxWithMethodAndAuth(method, bearer string) context.Context {
	ctx := grpc.NewContextWithServerTransportStream(context.Background(), &mockStream{method: method})
	if bearer != "" {
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("authorization", "bearer "+bearer))
	}
	return ctx
}

func ctxWithMethodAndRawAuth(method, authorization string) context.Context {
	ctx := grpc.NewContextWithServerTransportStream(context.Background(), &mockStream{method: method})
	if authorization != "" {
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("authorization", authorization))
	}
	return ctx
}

func uidFrom(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(utility.UIDContextKey).(string)
	return uid, ok && uid != ""
}

func TestAuthorAuth(t *testing.T) {
	t.Parallel()

	const method = "/game0.pb.Game0Service/Hi"
	author := &Author{
		client:        &mockAuthClient{uid: "uid-1"},
		unAuthMethods: map[string]struct{}{},
	}

	t.Run("missing bearer fails closed", func(t *testing.T) {
		t.Parallel()
		ctx, err := author.Auth(ctxWithMethodAndAuth(method, ""))
		if err == nil {
			t.Fatal("expected error")
		}
		if status.Code(err) != codes.Unauthenticated {
			t.Fatalf("code=%v want Unauthenticated", status.Code(err))
		}
		if _, ok := uidFrom(ctx); ok {
			t.Fatal("uid must not be set on failure")
		}
	})

	t.Run("malformed authorization fails closed", func(t *testing.T) {
		t.Parallel()
		ctx, err := author.Auth(ctxWithMethodAndRawAuth(method, "not-a-scheme-token"))
		if err == nil {
			t.Fatal("expected error")
		}
		if status.Code(err) != codes.Unauthenticated {
			t.Fatalf("code=%v want Unauthenticated", status.Code(err))
		}
		if _, ok := uidFrom(ctx); ok {
			t.Fatal("uid must not be set on failure")
		}
	})

	t.Run("validate failure fails closed", func(t *testing.T) {
		t.Parallel()
		mock := &mockAuthClient{err: status.Error(codes.Unauthenticated, "bad token")}
		a := &Author{
			client:        mock,
			unAuthMethods: map[string]struct{}{},
		}
		ctx, err := a.Auth(ctxWithMethodAndAuth(method, "tok-xyz"))
		if err == nil {
			t.Fatal("expected error")
		}
		if status.Code(err) != codes.Unauthenticated {
			t.Fatalf("code=%v want Unauthenticated", status.Code(err))
		}
		if mock.lastAccessTok != "tok-xyz" {
			t.Fatalf("token forwarded=%q want tok-xyz", mock.lastAccessTok)
		}
		if _, ok := uidFrom(ctx); ok {
			t.Fatal("uid must not be set on validate failure")
		}
	})

	t.Run("validate plain error still fails and forwards token", func(t *testing.T) {
		t.Parallel()
		mock := &mockAuthClient{err: errors.New("bad token")}
		a := &Author{
			client:        mock,
			unAuthMethods: map[string]struct{}{},
		}
		ctx, err := a.Auth(ctxWithMethodAndAuth(method, "tok-plain"))
		if err == nil {
			t.Fatal("expected error")
		}
		if mock.lastAccessTok != "tok-plain" {
			t.Fatalf("token forwarded=%q want tok-plain", mock.lastAccessTok)
		}
		if _, ok := uidFrom(ctx); ok {
			t.Fatal("uid must not be set on failure")
		}
	})

	t.Run("success sets uid and forwards token", func(t *testing.T) {
		t.Parallel()
		mock := &mockAuthClient{uid: "uid-1"}
		a := &Author{
			client:        mock,
			unAuthMethods: map[string]struct{}{},
		}
		ctx, err := a.Auth(ctxWithMethodAndAuth(method, "tok-ok"))
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if mock.lastAccessTok != "tok-ok" {
			t.Fatalf("token forwarded=%q want tok-ok", mock.lastAccessTok)
		}
		uid, ok := uidFrom(ctx)
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
		if _, ok := uidFrom(ctx); ok {
			t.Fatal("uid must not be set on unauth bypass")
		}
	})
}
