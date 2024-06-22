package bi

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFoundEventType = status.Error(codes.NotFound, "ErrNotFoundEventType")
	ErrInvalidProperties = status.Error(codes.InvalidArgument, "ErrInvalidProperties")
)
