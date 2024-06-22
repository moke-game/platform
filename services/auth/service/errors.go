package public

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrGeneralFailure       = status.Error(codes.Internal, "ErrGeneralFailure")
	ErrGenerateJwtFailure   = status.Error(codes.Internal, "ErrGenerateJwtFailure")
	ErrClientParamFailure   = status.Error(codes.Internal, "ErrClientParamFailure")
	ErrParseJwtTokenFailure = status.Error(codes.Internal, "ErrParseJwtTokenFailure")
	ErrPermissionDenied     = status.Error(codes.PermissionDenied, "PermissionDenied")
)
