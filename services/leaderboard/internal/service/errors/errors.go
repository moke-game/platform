package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoMetaData     = status.Error(codes.Internal, "ErrNoMetaData")
	ErrExpireFailed   = status.Error(codes.Internal, "ErrExpireFailed")
	ErrMaxNum         = status.Error(codes.Internal, "ErrMaxNum")
	ErrGeneralFailure = status.Error(codes.Internal, "ErrGeneralFailure")
)
