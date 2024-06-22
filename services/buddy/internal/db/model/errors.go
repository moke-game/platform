package model

import "errors"

var (
	ErrLoadBuddyQueueFailed = errors.New("ErrLoadBuddyQueueFailed")
	ErrBuddyQueueNotFound   = errors.New("ErrBuddyQueueNotFound")
	ErrNoDataChange         = errors.New("ErrNoDataChange")
)
