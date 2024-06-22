package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoMetaData     = status.Error(codes.PermissionDenied, "ErrNoMetaData")
	ErrGeneralFailure = status.Error(codes.Internal, "ErrGeneralFailure")
	ErrPartyNotFound  = status.Error(codes.NotFound, "ErrPartyNotFound")
	ErrNotOwner       = status.Error(codes.PermissionDenied, "ErrNotOwner")
	ErrIllegal        = status.Error(codes.PermissionDenied, "ErrIllegal")
	ErrPartyFull      = status.Error(codes.PermissionDenied, "ErrPartyFull")
	ErrHasParty       = status.Error(codes.Internal, "ErrHasParty")
)
