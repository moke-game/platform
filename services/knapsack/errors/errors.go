package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoMetaData     = status.Error(codes.Internal, "ErrNoMetaData")
	ErrNotFound       = status.Error(codes.NotFound, "ErrNotFound")
	ErrNotEnough      = status.Error(codes.Internal, "ErrNotEnough")
	ErrGeneralFailure = status.Error(codes.Internal, "ErrGeneralFailure")
)
