package db

import "github.com/gstones/moke-kit/orm/nosql/key"

func makePartyKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("party", "manager", id)
}

func makePartyMemberKey(partyId string) (key.Key, error) {
	return key.NewKeyFromParts("party", "members", partyId)
}

func makeUid2PidKey() (key.Key, error) {
	return key.NewKeyFromParts("party", "uid2pid")
}

func makeInviteKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("party", "invite", id)
}

func makePartyIdKey() (key.Key, error) {
	return key.NewKeyFromParts("party", "generate", "id")
}
