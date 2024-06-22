package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrParamInvalid = status.Error(codes.InvalidArgument, "ErrParamInvalid")
)
