package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoMetaData        = status.Error(codes.InvalidArgument, "ErrNoMetaData")
	ErrParamsInvalid     = status.Error(codes.InvalidArgument, "ErrParamsInvalid")
	ErrSaveMailFailed    = status.Error(codes.Internal, "ErrSaveMailFailed")
	ErrPublishMailFailed = status.Error(codes.Internal, "ErrPublishMailFailed")
)
