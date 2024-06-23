package data

import (
	"sort"

	pb "github.com/moke-game/platform/api/gen/buddy"
)

type MemberType int32

const (
	// Already a buddy.
	MemberTypeBuddy MemberType = 0
	// Need to be identified by self.
	MemberTypeInviter MemberType = 1
)

type Buddy struct {
	// The UID for a specific buddy.
	UID string
	// Timestamp (UTC) when request add buddy was send.
	ActTime int64
	// The Remark for buddy.
	Remark string
	// Reward  num from buddy.
	ReceiveRewardNum int32
	// FriendValue int32
	FriendValue int32
	// isFavorite
	IsFavorite bool
}

type Inviter struct {
	// The UID for a specific buddy.
	UID string
	// Timestamp (UTC) when request add buddy was send.
	ReqTime int64
	// Request add buddy text.
	ReqInfo string
}

func (i *Inviter) ToProto() *pb.Inviter {
	return &pb.Inviter{
		Uid:     i.UID,
		ReqInfo: i.ReqInfo,
		ReqTime: i.ReqTime,
	}
}

func (b *Buddy) ToProto() *pb.Buddy {
	return &pb.Buddy{
		Uid:           b.UID,
		ReceiveReward: b.ReceiveRewardNum,
		IsFavorite:    b.IsFavorite,
		Remark:        b.Remark,
		ActTime:       b.ActTime,
	}
}

func (b *Buddy) IncrFriendValue(value int32) {
	b.FriendValue += value
}

func (b *Buddy) ReceiveReward() {
	b.ReceiveRewardNum++
}

func (b *Buddy) Favorite(isFavor bool) {
	b.IsFavorite = isFavor
}

func (b *Buddy) UnFavorite() {
	b.IsFavorite = false
}
func NewInviter(uid string, text string, time int64) *Inviter {
	return &Inviter{
		UID:     uid,
		ReqTime: time,
		ReqInfo: text,
	}
}

func NewBuddy(uid string, remark string, time int64) *Buddy {
	return &Buddy{
		UID:     uid,
		ActTime: time,
		Remark:  remark,
	}
}

type BuddySettings struct {
	AllowToBeAdded bool
}

type BuddyQueue struct {
	UID      string
	Nickname string

	Buddies         map[string]*Buddy
	Inviters        map[string]*Inviter
	InviterSends    map[string]*Inviter
	BlockedProfiles map[string]*BlockedProfile
	RecentMet       []*RecentProfile
	Settings        *BuddySettings
}

func NewBuddyQueue(uid string) *BuddyQueue {
	return &BuddyQueue{
		UID:             uid,
		Buddies:         make(map[string]*Buddy),
		BlockedProfiles: make(map[string]*BlockedProfile),
		Inviters:        make(map[string]*Inviter),
		InviterSends:    make(map[string]*Inviter),
		Settings: &BuddySettings{
			AllowToBeAdded: true,
		},
	}
}
func (bq *BuddyQueue) Clear() {
	bq.Buddies = make(map[string]*Buddy)
	bq.Inviters = make(map[string]*Inviter)
	bq.InviterSends = make(map[string]*Inviter)
	bq.BlockedProfiles = make(map[string]*BlockedProfile)
	bq.Settings = &BuddySettings{
		AllowToBeAdded: true,
	}
}

// AddBuddy Add a new buddy instance to the queue.
func (bq *BuddyQueue) AddBuddy(f *Buddy) {
	bq.Buddies[f.UID] = f
}

func (bq *BuddyQueue) AddInviter(f *Inviter) {
	bq.Inviters[f.UID] = f
}

func (bq *BuddyQueue) GetInviterNum() int32 {
	return int32(len(bq.Inviters))
}
func (bq *BuddyQueue) GetInviterIds() []string {
	res := make([]string, 0)
	for k := range bq.Inviters {
		res = append(res, k)
	}
	return res
}
func (bq *BuddyQueue) GetInviters() []*Inviter {
	res := make([]*Inviter, 0)
	for _, v := range bq.Inviters {
		res = append(res, v)
	}
	return res
}

func (bq *BuddyQueue) GetSortedInviters() []string {
	inviters := bq.GetInviters()
	sort.Slice(inviters, func(i, j int) bool {
		return inviters[i].ReqTime > inviters[j].ReqTime
	})
	res := make([]string, 0)
	for _, v := range inviters {
		res = append(res, v.UID)
	}
	return res
}

func (bq *BuddyQueue) GetInviter(uid string) *Inviter {
	if v, ok := bq.Inviters[uid]; ok {
		return v
	}
	return nil
}

func (bq *BuddyQueue) IsInvited(id string) bool {
	if _, ok := bq.Inviters[id]; ok {
		return true
	}
	return false
}

func (bq *BuddyQueue) AddInviteSend(uid string, in *Inviter) {
	bq.InviterSends[uid] = in
}

func (bq *BuddyQueue) RemoveInviteSend(uid string) {
	delete(bq.InviterSends, uid)
}

func (bq *BuddyQueue) IsAlreadySendInvited(id string) bool {
	if _, ok := bq.InviterSends[id]; ok {
		return true
	}
	return false
}

func (bq *BuddyQueue) GetBuddy(uid string) *Buddy {
	if v, ok := bq.Buddies[uid]; ok {
		return v
	}
	return nil
}

func (bq *BuddyQueue) RemoveInviter(uid string) {
	delete(bq.Inviters, uid)
}

func (bq *BuddyQueue) Favorite(isFavor bool, ids ...string) {
	for _, v := range ids {
		if b, ok := bq.Buddies[v]; ok {
			b.Favorite(isFavor)
		}
	}
}

func (bq *BuddyQueue) CollectRewardNum(ids []string) int32 {
	res := int32(0)
	for _, v := range ids {
		if b := bq.GetBuddy(v); b != nil {
			res += b.ReceiveRewardNum
			b.ReceiveRewardNum = 0
		}
	}
	return res
}

func (bq *BuddyQueue) ClearRewardNum() (int32, []string) {
	res := int32(0)
	ids := make([]string, 0)
	for _, v := range bq.Buddies {
		if v.ReceiveRewardNum > 0 {
			res += v.ReceiveRewardNum
			v.ReceiveRewardNum = 0
			ids = append(ids, v.UID)
		}
	}
	return res, ids
}

func (bq *BuddyQueue) AddFriendValue(value int32, ids ...string) {
	for _, v := range ids {
		if b := bq.GetBuddy(v); b != nil {
			b.IncrFriendValue(value)
		}
	}
}

func (bq *BuddyQueue) Delete(uid string) {
	delete(bq.Buddies, uid)
}
func (bq *BuddyQueue) DeleteBuddies(ids ...string) {
	for _, v := range ids {
		delete(bq.Buddies, v)
	}
}

func (bq *BuddyQueue) GetMemberCounts() int32 {
	return int32(len(bq.Buddies))
}

func (bq *BuddyQueue) IsContains(uid string) bool {
	_, ok := bq.Buddies[uid]
	return ok
}
func (bq *BuddyQueue) IsContainsProfiles(ids ...string) bool {
	for _, v := range ids {
		if !bq.IsContains(v) {
			return false
		}
	}
	return true
}

func (bq *BuddyQueue) IsContainsInviter(uid string) bool {
	_, ok := bq.Inviters[uid]
	return ok
}

func (bq *BuddyQueue) UpdateRemark(uid string, remark string) {
	if b := bq.GetBuddy(uid); b != nil {
		b.Remark = remark
		return
	}
}

func (bq *BuddyQueue) UpdateSettings(allowToBeAdded bool) {
	bq.Settings.AllowToBeAdded = allowToBeAdded
}
