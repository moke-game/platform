package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidParameter              = status.Error(codes.InvalidArgument, "ErrInvalidParameter")
	ErrGeneralFailure                = status.Error(codes.Internal, "ErrGeneralFailure")
	ErrBuddyAlreadyAdded             = status.Error(codes.AlreadyExists, "ErrBuddyAlreadyAdded")
	ErrBuddyAlreadyRequested         = status.Error(codes.AlreadyExists, "ErrBuddyAlreadyRequested")
	ErrBuddyAlreadyInYourRequestList = status.Error(codes.AlreadyExists, "ErrBuddyAlreadyInYourRequestList")
	ErrNoMetaData                    = status.Error(codes.PermissionDenied, "ErrNoMetaData")
	ErrNotAllowed                    = status.Error(codes.PermissionDenied, "ErrNotAllowed")
	ErrBuddiesNotFound               = status.Error(codes.NotFound, "ErrBuddiesNotFound")
	ErrInviterNotFound               = status.Error(codes.NotFound, "ErrInviterNotFound")
	ErrSelfBuddiesTopLimit           = status.Error(codes.Internal, "ErrSelfBuddiesTopLimit")
	ErrSelfInviterTopLimit           = status.Error(codes.Internal, "ErrSelfInviterTopLimit")
	ErrTargetInviterTopLimit         = status.Error(codes.Internal, "ErrTargetInviterTopLimit")
	ErrTargetBuddiesTopLimit         = status.Error(codes.Internal, "ErrTargetBuddiesTopLimit")
	ErrCanNotAddSelf                 = status.Error(codes.Internal, "ErrCanNotAddSelf")
	ErrInTargetBlockedList           = status.Error(codes.Internal, "ErrInTargetBlockedList")
	ErrInSelfBlockedList             = status.Error(codes.Internal, "ErrInSelfBlockedList")
	ErrBlockedNumExceed              = status.Error(codes.Internal, "ErrBlockedNumExceed")
	ErrDBErr                         = status.Error(codes.Internal, "ErrDBErr")
)
