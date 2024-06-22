package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoMetaData              = status.Error(codes.PermissionDenied, "ErrNoMetaData")
	ErrNoSuchChatDestination   = status.Error(codes.NotFound, "ErrNoSuchDestination")
	ErrSubscribeChatFailed     = status.Error(codes.Unavailable, "ErrSubscribeChatFailed")
	ErrUnSubscribeParamFailed  = status.Error(codes.Unavailable, "ErrUnSubscribeParamFailed")
	ErrSubscribeRoomIDFailed   = status.Error(codes.Unavailable, "ErrSubscribeRoomIDFailed")
	ErrUnSubscribeRoomIDFailed = status.Error(codes.Unavailable, "ErrUnSubscribeRoomIDFailed")
	ErrChatRoomIDFailed        = status.Error(codes.Unavailable, "ErrChatRoomIDFailed")
	ErrPublishChatFailed       = status.Error(codes.Unavailable, "ErrPublishChatFailed")
	ErrSameClientAliveFailed   = status.Error(codes.Unavailable, "ErrSameClientAliveFailed")
	ErrSubscribeParamFailed    = status.Error(codes.Unavailable, "ErrSubscribeParamFailed")
	ErrChatTypeFailed          = status.Error(codes.Unavailable, "ErrChatTypeFailed")
	ErrClientProfileIdNotSame  = status.Error(codes.Unavailable, "ErrClientProfileIdNotSame")
	ErrDestinationFailure      = status.Error(codes.Unavailable, "ErrDestinationFailure")
	ErrChatMessageFailure      = status.Error(codes.Unavailable, "ErrChatMessageFailure")
	ErrChatIntervalFailed      = status.Error(codes.Unavailable, "ErrChatIntervalFailed")
	ErrNoSuchClient            = status.Error(codes.Internal, "ErrNoSuchClient")
	ErrBadRequest              = status.Error(codes.InvalidArgument, "ErrBadRequest")
	ErrGeneralFailure          = status.Error(codes.Internal, "ErrGeneralFailure")
	ErrConnectionFailure       = status.Error(codes.Internal, "ErrConnectionFailure")
	ErrReSubscribeChatChannels = status.Error(codes.Unavailable, "ErrReSubscribeChatChannels")
	ErrRequireSubscribeFirst   = status.Error(codes.Unavailable, "ErrRequireSubscribeFirst")
)
