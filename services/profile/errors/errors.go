package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoMetaData      = status.Error(codes.Internal, "ErrNoMetaData")
	ErrNotFound        = status.Error(codes.NotFound, "ErrNotFound")
	ErrGeneralFailure  = status.Error(codes.Internal, "ErrGeneralFailure")
	ErrUpdateFailure   = status.Error(codes.Internal, "ErrUpdateFailure")
	ErrLoadFailure     = status.Error(codes.Internal, "ErrLoadFailure")
	ErrAlreadyExists   = status.Error(codes.AlreadyExists, "ErrAlreadyExists")
	ErrInvalidArgument = status.Error(codes.InvalidArgument, "ErrInvalidArgument")
)
