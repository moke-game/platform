package db

import (
	"strconv"

	"github.com/gstones/moke-kit/orm/nosql/key"
)

const (
	MailTheme           = "mail"
	MailPublicIndex     = "index"
	MailPublicKeyPrefix = "public"
)

func makeMailKey(profileId string) (key.Key, error) {
	return key.NewKeyFromParts(MailTheme, profileId)
}

func makeMailPublicListKey(channel string) (key.Key, error) {
	if channel == "" {
		channel = "0"
	}
	return key.NewKeyFromParts(MailPublicKeyPrefix, channel)
}

func makeMailPublicIndexKey(profileId string, channel string) (key.Key, error) {
	if channel == "" {
		channel = "0"
	}
	return key.NewKeyFromParts(MailTheme, channel, MailPublicIndex, profileId)
}

func makeFieldMailKey(uid int64) (key.Key, error) {
	idStr := strconv.FormatInt(uid, 10)
	return key.NewKeyFromParts(MailTheme, idStr)
}
