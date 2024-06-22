package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoMetaData     = status.Error(codes.PermissionDenied, "ErrNoMetaData")
	ErrGeneralFailure = status.Error(codes.Internal, "ErrGeneralFailure")
)
