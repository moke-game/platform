package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrGeneralFailure        = status.Error(codes.Internal, "ErrGeneralFailure")
	ErrBuddyAlreadyAdded     = status.Error(codes.AlreadyExists, "ErrBuddyAlreadyAdded")
	ErrBuddyAlreadyRequested = status.Error(codes.AlreadyExists, "ErrBuddyAlreadyRequested")
	ErrNoMetaData            = status.Error(codes.PermissionDenied, "ErrNoMetaData")
	ErrBuddiesNotFound       = status.Error(codes.NotFound, "ErrBuddiesNotFound")
	ErrInviterNotFound       = status.Error(codes.NotFound, "ErrInviterNotFound")
	ErrSelfBuddiesTopLimit   = status.Error(codes.Internal, "ErrSelfBuddiesTopLimit")
	ErrTargetInviterTopLimit = status.Error(codes.Internal, "ErrTargetInviterTopLimit")
	ErrTargetBuddiesTopLimit = status.Error(codes.Internal, "ErrTargetBuddiesTopLimit")
	ErrCanNotAddSelf         = status.Error(codes.Internal, "ErrCanNotAddSelf")
	ErrInTargetBlockedList   = status.Error(codes.Internal, "ErrInTargetBlockedList")
	ErrInSelfBlockedList     = status.Error(codes.Internal, "ErrInSelfBlockedList")
	ErrDBErr                 = status.Error(codes.Internal, "ErrDBErr")
)
