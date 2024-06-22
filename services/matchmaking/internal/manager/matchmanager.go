package manager

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/google/uuid"
	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	pb2 "github.com/gstones/platform/api/gen/auth"
	pb "github.com/gstones/platform/api/gen/matchmaking"
	"github.com/gstones/platform/services/matchmaking/internal/agones"
	"github.com/gstones/platform/services/matchmaking/internal/db"
	"github.com/gstones/platform/services/matchmaking/internal/db/model"
	"github.com/gstones/platform/services/matchmaking/internal/utils"
	"github.com/gstones/platform/services/matchmaking/pkg/module/data"
)

const (
	// MaxLevel .
	MaxLevel = 3
)

var (
	matchManager *MatchManager
	nextRetry    = []int64{5, 10, 20}
)

func NewMatchManager(
	db *db.Database,
	l *zap.Logger,
	mq miface.MessageQueue,
	aClient pb2.AuthServiceClient,
	allocator *agones.Allocator,
) *MatchManager {
	matchManager = &MatchManager{
		db:          db,
		mq:          mq,
		playMap:     sync.Map{},
		pveMap:      sync.Map{},
		playerDic:   sync.Map{},
		retryMap:    sync.Map{},
		retryIdsMap: sync.Map{},
		logger:      l,
		aClient:     aClient,
		allocator:   allocator,
	}
	return matchManager
}

type MatchManager struct {
	db          *db.Database
	mq          miface.MessageQueue
	playMap     sync.Map //玩法列表
	pveMap      sync.Map //PVE玩法列表
	guidMap     sync.Map //新手引导战斗玩法列表
	playerDic   sync.Map //key:玩家id value:redis key
	retryMap    sync.Map //匹配成功没有成功分配房间的重试队列 map[string][]*model.MatchData
	retryIdsMap sync.Map //匹配成功没有成功分配房间的重试队列 key:玩家id value:房间id
	logger      *zap.Logger
	aClient     pb2.AuthServiceClient
	allocator   *agones.Allocator
}

func GetGlobalMatchManager() *MatchManager {
	return matchManager
}

func (m *MatchManager) Update(_ int32) {
	m.playMap.Range(func(key, value interface{}) bool {
		playId := key.(string)
		for i := 0; i < MaxLevel; i++ {
			m.checkWaitLevel(string(db.Match_Type_Single), playId, i)
			m.checkWaitLevel(string(db.Match_Type_Team), playId, i)
			m.checkMatchLevel(string(db.Match_Type_Single), playId, i)
			m.checkMatchLevel(string(db.Match_Type_Team), playId, i)
			m.checkPEVWaitLevel(string(db.Match_Type_PVE), playId, i)
		}
		return true
	})
	m.pveMap.Range(func(key, value interface{}) bool {
		playId := key.(string)
		for i := 0; i < MaxLevel; i++ {
			m.checkPEVWaitLevel(string(db.Match_Type_PVE), playId, i)
		}
		return true
	})
	m.guidMap.Range(func(key, value interface{}) bool {
		playId := key.(string)
		for i := 0; i < MaxLevel; i++ {
			m.checkGuidWaitLevel(string(db.Match_Type_Guid), playId, i)
		}
		return true
	})

}

func (m *MatchManager) UpdateRetry(_ int32) {
	nowTime := time.Now().UTC().Unix()
	m.retryMap.Range(func(key, value interface{}) bool {
		roomId := key.(string)
		match := value.(*model.MatchRetry)
		if nowTime < match.NextRetryTime {
			return true
		}
		if ok := m.onMatchSuccess(match.MatchData, true); ok {
			m.delRetryMatchByRoomId(roomId)
		} else {
			match.RetryCount++
			if len(nextRetry) <= match.RetryCount {
				m.delRetryMatchByRoomId(roomId)
				ids := make([]string, 0)
				for _, datum := range match.MatchData {
					for _, playerData := range datum.Members {
						ids = append(ids, playerData.Uid)
					}
				}
				m.pubCancelMatchMsg(ids)
				return true
			}
			match.NextRetryTime = nowTime + nextRetry[match.RetryCount]
		}
		return true
	})
}

func (m *MatchManager) checkMatchLevel(matchType, playId string, queueLevel int) map[*model.MatchData]int {
	queueLevelStr := strconv.Itoa(queueLevel)
	dataMap, err := m.db.GetAllMatchData(matchType, queueLevelStr, playId)
	if err != nil {
		m.logger.Error("matchmanager checkMatchLevel err", zap.String("matchtype", matchType), zap.Int("queuelevel", queueLevel), zap.Error(err))
		return nil
	}
	result := make(map[*model.MatchData]int)
	now := time.Now().Unix()
	for key, md := range dataMap {
		sub := now - md.MatchTime
		level := checkQueueLevel(sub)
		if level > queueLevel {
			//删除队列
			err := m.db.DeleteData(matchType, string(db.Queue_Type_Match), queueLevelStr, playId, key)
			if err != nil {
				continue
			}
			result[md] = level
		}
	}
	for matchData, i := range result {
		m.checkMatchmaking(matchData, matchType, i)
	}
	return result
}

func (m *MatchManager) checkWaitLevel(matchType, playId string, queueLevel int) {
	queueLevelStr := strconv.Itoa(queueLevel)
	dataMap, err := m.db.GetAllWaitData(matchType, queueLevelStr, playId)
	if err != nil {
		m.logger.Error("matchmanager checkWaitLevel err", zap.String("matchtype", matchType), zap.Int("queuelevel", queueLevel), zap.Error(err))
		return
	}
	result := make(map[*model.MatchData]int)
	now := time.Now().Unix()
	for key, md := range dataMap {
		sub := now - md.MatchTime
		level := checkQueueLevel(sub)
		if level > queueLevel {
			//删除队列
			err := m.db.DeleteData(matchType, string(db.Queue_Type_Wait), queueLevelStr, playId, key)
			if err != nil {
				continue
			}
			result[md] = level
		}
	}
	for matchData, level := range result {
		m.checkTeammate(matchData, matchType, level)
	}
}

func (m *MatchManager) checkPEVWaitLevel(matchType, playId string, queueLevel int) {
	queueLevelStr := strconv.Itoa(queueLevel)
	dataMap, err := m.db.GetAllWaitData(matchType, queueLevelStr, playId)
	if err != nil {
		m.logger.Error("matchmanager checkPEVWaitLevel err", zap.String("matchtype", matchType), zap.Int("queuelevel", queueLevel), zap.Error(err))
		return
	}
	result := make(map[*model.MatchData]int)
	now := time.Now().Unix()
	for key, md := range dataMap {
		sub := now - md.MatchTime
		level := checkQueueLevel(sub)
		if level > queueLevel {
			//删除队列
			err := m.db.DeleteData(matchType, string(db.Queue_Type_Wait), queueLevelStr, playId, key)
			if err != nil {
				continue
			}
			result[md] = level
		}
	}
	for matchData, level := range result {
		m.checkPVETeammate(matchData, matchType, level)
	}
}

func (m *MatchManager) checkGuidWaitLevel(matchType, playId string, queueLevel int) {
	queueLevelStr := strconv.Itoa(queueLevel)
	dataMap, err := m.db.GetAllWaitData(matchType, queueLevelStr, playId)
	if err != nil {
		m.logger.Error("matchmanager checkGuidWaitLevel err", zap.String("matchtype", matchType), zap.Int("queuelevel", queueLevel), zap.Error(err))
		return
	}
	result := make(map[*model.MatchData]int)
	now := time.Now().Unix()
	for key, md := range dataMap {
		sub := now - md.MatchTime
		level := checkQueueLevel(sub)
		if level > queueLevel {
			//删除队列
			err := m.db.DeleteData(matchType, string(db.Queue_Type_Wait), queueLevelStr, playId, key)
			if err != nil {
				continue
			}
			result[md] = level
		}
	}
	for matchData, level := range result {
		//已经超时
		if level >= MaxLevel {
			robotMatchData := &model.MatchData{
				Members:   make(map[string]*model.PlayerData),
				GroupSize: matchData.GroupSize,
				Score:     matchData.Score,
			}
			list := []*model.MatchData{matchData, robotMatchData}
			m.onMatchSuccess(list, false)
			continue
		}
		//保存数据到新的队列
		err := m.db.PushData(matchType, string(db.Queue_Type_Wait), strconv.Itoa(queueLevel), playId, matchData)
		if err != nil {
			return
		}
		m.updatePlayerDic(matchData, matchType, string(db.Queue_Type_Wait), strconv.Itoa(queueLevel), playId)
	}
}

func (m *MatchManager) checkMatchmaking(matchData *model.MatchData, matchType string, queueLevel int) {
	//检测匹配队列里是否有符合条件的队伍
	playIdStr := strconv.Itoa(int(matchData.PlayId))
	for i := 2; i >= 0; i-- {
		queueLevelStr := strconv.Itoa(i)
		dataMap, err := m.db.GetAllMatchData(matchType, queueLevelStr, playIdStr)
		if err != nil {
			m.logger.Error(
				"checkMatchmaking GetAllMatchData err",
				zap.String("matchtype", matchType),
				zap.Int("queuelevel", queueLevel),
				zap.Error(err),
			)
			return
		}
		maxSubLevel := i
		//已经超时
		if queueLevel >= MaxLevel {
			maxSubLevel = 99
		} else {
			if maxSubLevel < queueLevel {
				maxSubLevel = queueLevel
			}
		}

		for key, d := range dataMap {
			subLevel := int(math.Abs(float64(matchData.Score - d.Score)))
			if d.GroupSize != 3 && subLevel > maxSubLevel {
				continue
			}
			//匹配成功
			//先检测是否存在
			exit, _ := m.db.HaseData(matchType, string(db.Queue_Type_Match), queueLevelStr, playIdStr, key)
			if !exit {
				continue
			}
			//从队列中移除
			er := m.db.DeleteData(matchType, string(db.Queue_Type_Match), queueLevelStr, playIdStr, key)
			if er != nil {
				m.logger.Error("checkMatchmaking DeleteData err", zap.Error(er))
				continue
			}
			list := []*model.MatchData{matchData, d}
			m.onMatchSuccess(list, false)
			return
		}
	}
	//检测等待队列里是否有符合条件的队伍
	for i := 2; i >= 0; i-- {
		queueLevelStr := strconv.Itoa(i)
		dataMap, err := m.db.GetAllWaitData(matchType, queueLevelStr, playIdStr)
		if err != nil {
			m.logger.Error("checkMatchmaking GetAllWaitData err", zap.String("matchtype", matchType), zap.Int("queuelevel", queueLevel), zap.Error(err))
			return
		}
		maxSubLevel := i
		//已经超时
		if queueLevel >= MaxLevel {
			maxSubLevel = 99
		} else {
			if maxSubLevel < queueLevel {
				maxSubLevel = queueLevel
			}
		}

		for key, d := range dataMap {
			subLevel := int(math.Abs(float64(matchData.Score - d.Score)))
			if subLevel > maxSubLevel {
				continue
			}
			//匹配成功
			//先检测是否存在
			exit, _ := m.db.HaseData(matchType, string(db.Queue_Type_Wait), queueLevelStr, playIdStr, key)
			if !exit {
				continue
			}
			//从队列中移除
			er := m.db.DeleteData(matchType, string(db.Queue_Type_Wait), queueLevelStr, playIdStr, key)
			if er != nil {
				m.logger.Error("checkMatchmaking DeleteData err", zap.Error(er))
				continue
			}
			list := []*model.MatchData{matchData, d}
			m.onMatchSuccess(list, false)
			return
		}
	}
	//超时匹配 直接匹配成功
	if queueLevel >= MaxLevel {
		robotMatchData := &model.MatchData{
			Members:   make(map[string]*model.PlayerData),
			GroupSize: matchData.GroupSize,
			Score:     matchData.Score,
		}
		list := []*model.MatchData{matchData, robotMatchData}
		m.onMatchSuccess(list, false)
		return
	}
	queueLevelStr := strconv.Itoa(queueLevel)
	//保存数据到新的队列
	err := m.db.PushData(matchType, string(db.Queue_Type_Match), queueLevelStr, playIdStr, matchData)
	if err != nil {
		m.logger.Error("checkMatchmaking PushData  err", zap.Error(err))
		return
	}
	m.updatePlayerDic(matchData, matchType, string(db.Queue_Type_Match), queueLevelStr, playIdStr)
}

func (m *MatchManager) checkTeammate(matchData *model.MatchData, matchType string, queueLevel int) {
	size := int(matchData.GroupSize)
	playIdStr := strconv.Itoa(int(matchData.PlayId))
	//已经满员或者超时 直接进入匹配
	if len(matchData.Members) == size || queueLevel >= MaxLevel {
		m.checkMatchmaking(matchData, matchType, queueLevel)
		return
	}
	for i := 2; i >= 0; i-- {
		queueStr := strconv.Itoa(i)
		dataMap, err := m.db.GetAllWaitData(matchType, queueStr, playIdStr)
		if err != nil {
			m.logger.Error("matchmanager checkTeammate err", zap.String("matchtype", matchType), zap.String("queuelevel", queueStr), zap.Error(err))
			continue
		}
		maxSubLevel := i
		if maxSubLevel < queueLevel {
			maxSubLevel = queueLevel
		}
		for key, d := range dataMap {
			subLevel := int(math.Abs(float64(matchData.Score - d.Score)))
			if subLevel > i {
				continue
			}
			if len(d.Members)+len(matchData.Members) > size {
				continue
			}
			//合并两个队伍
			d.AppendMember(matchData)
			//更新分数
			d.UpdateScore()
			//先检测是否存在
			exit, _ := m.db.HaseData(matchType, string(db.Queue_Type_Wait), queueStr, playIdStr, key)
			if !exit {
				continue
			}
			//已经满员
			if len(d.Members) == size {
				//从队列中移除
				err = m.db.DeleteData(matchType, string(db.Queue_Type_Wait), queueStr, playIdStr, key)
				if err != nil {
					continue
				}
				m.checkMatchmaking(d, matchType, i)
				return
			} else {
				//更新队列中的数据
				err = m.db.PushData(matchType, string(db.Queue_Type_Wait), queueStr, playIdStr, d)
				if err != nil {
					continue
				}
				m.updatePlayerDic(matchData, matchType, string(db.Queue_Type_Wait), queueStr, playIdStr)
				return
			}

		}
	}
	//保存数据到新的队列
	err := m.db.PushData(matchType, string(db.Queue_Type_Wait), strconv.Itoa(queueLevel), playIdStr, matchData)
	if err != nil {
		return
	}
	m.updatePlayerDic(matchData, matchType, string(db.Queue_Type_Wait), strconv.Itoa(queueLevel), playIdStr)
}

func (m *MatchManager) checkPVETeammate(matchData *model.MatchData, matchType string, queueLevel int) {
	size := int(matchData.GroupSize)
	playIdStr := strconv.Itoa(int(matchData.PlayId))
	//已经满员或者超时
	if len(matchData.Members) == size || queueLevel >= MaxLevel {
		m.onMatchSuccess([]*model.MatchData{matchData}, false)
		return
	}
	for i := 2; i >= 0; i-- {
		queueStr := strconv.Itoa(i)
		dataMap, err := m.db.GetAllWaitData(matchType, queueStr, playIdStr)
		if err != nil {
			m.logger.Error("matchmanager checkPVETeammate err", zap.String("matchtype", matchType), zap.String("queuelevel", queueStr), zap.Error(err))
			continue
		}
		maxSubLevel := i
		if maxSubLevel < queueLevel {
			maxSubLevel = queueLevel
		}
		for key, d := range dataMap {
			subLevel := int(math.Abs(float64(matchData.Score - d.Score)))
			if subLevel > i {
				continue
			}
			if len(d.Members)+len(matchData.Members) > size {
				continue
			}
			//合并两个队伍
			d.AppendMember(matchData)
			//更新分数
			d.UpdateScore()
			//先检测是否存在
			exit, _ := m.db.HaseData(matchType, string(db.Queue_Type_Wait), queueStr, playIdStr, key)
			if !exit {
				continue
			}
			//已经满员
			if len(d.Members) == size {
				//从队列中移除
				err = m.db.DeleteData(matchType, string(db.Queue_Type_Wait), queueStr, playIdStr, key)
				if err != nil {
					continue
				}
				m.onMatchSuccess([]*model.MatchData{matchData}, false)
				return
			} else {
				//更新队列中的数据
				err = m.db.PushData(matchType, string(db.Queue_Type_Wait), queueStr, playIdStr, d)
				if err != nil {
					continue
				}
				m.updatePlayerDic(matchData, matchType, string(db.Queue_Type_Wait), queueStr, playIdStr)
				return
			}
		}
	}
	//保存数据到新的队列
	err := m.db.PushData(matchType, string(db.Queue_Type_Wait), strconv.Itoa(queueLevel), playIdStr, matchData)
	if err != nil {
		return
	}
	m.updatePlayerDic(matchData, matchType, string(db.Queue_Type_Wait), strconv.Itoa(queueLevel), playIdStr)
}

func (m *MatchManager) onMatchSuccess(matchData []*model.MatchData, isRetry bool) bool {
	members := make(map[string]int32)
	robots := make(map[int32][]int32)
	players := make(map[string]data.PlayerData)
	playId := matchData[0].PlayId
	mapIdArr := matchData[0].MapId
	uidsAllocate := make([]string, 0)
	playerId := ""
	for index, datum := range matchData {
		camp := int32(index + 1)
		for _, playerData := range datum.Members {
			members[playerData.Uid] = camp
			uidsAllocate = append(uidsAllocate, playerData.Uid)
			player := data.PlayerData{
				Uid:          playerData.Uid,
				Nickname:     playerData.Nickname,
				Avatar:       playerData.Avatar,
				HeroId:       playerData.HeroId,
				SkinId:       playerData.SkinId,
				HeroLevel:    playerData.HeroLevel,
				Attribute:    playerData.Attribute,
				PetProfileId: playerData.PetProfileId,
				HeroCups:     playerData.HeroCups,
				PetSkill:     playerData.PetSkill,
				IsAgain:      playerData.IsAgain,
			}
			players[player.Uid] = player
			if len(playerId) == 0 {
				playerId = playerData.Uid
			}
		}
		//人数不够补充机器人
		loop := int(datum.GroupSize) - len(datum.Members)
		if loop > 0 {
			rt := make([]int32, loop)
			for i := 0; i < loop; i++ {
				rt[i] = datum.Score
			}
			robots[camp] = rt
		}
	}
	roomId := uuid.New().String()
	addr := ""
	mapId := int32(0)
	isFirstEnter := true
	var err error
	// TODO @GuoLei临时大厅地址
	if playId == 0 {
		playerUid := ""
		for _, member := range matchData[0].Members {
			playerUid = member.Uid
		}
		//需要检测是否断线重连
		battleRoomData, err := m.db.GetBattleRoom(playerUid)
		if err == nil && battleRoomData != nil {
			nowTime := time.Now().UTC().Unix()
			validTim := (nowTime - battleRoomData.BeginTime) < battleRoomData.PlayTime
			if validTim && m.validRoom(battleRoomData.RoomId) {
				roomId = battleRoomData.RoomId
				addr = battleRoomData.Addr
				playId = battleRoomData.PlayId
				mapId = battleRoomData.MapId
				isFirstEnter = false
			} else {
				roomId = "10000"
			}
		} else {
			roomId = "10000"
		}
	}
	if len(addr) == 0 {
		addr, err = m.allocator.Allocate(playId, uidsAllocate...)
		if err != nil {
			m.logger.Error("matchmaking allocateBattle err", zap.Any("error", err))
			if !isRetry {
				m.addRetryMatch(roomId, matchData)
			}
			return false
		}
	}
	var mapData *model.PlayerMapData
	if mapId == 0 {
		mapCount := len(mapIdArr)
		mapIdList := make([]int32, 0)
		if mapCount > 1 {
			mapData, err = m.db.GetBattleMap(playerId, playId)
			if err != nil {
				m.logger.Error("matchmaking GetBattleMap error", zap.Error(err))
			}
			if mapData == nil {
				mapData = &model.PlayerMapData{
					Uid: playerId,
				}
			}
			for _, id := range mapIdArr {
				if id == mapData.MapId && mapData.MapCount >= 2 {
					continue
				}
				mapIdList = append(mapIdList, id)
			}
		} else {
			mapIdList = append(mapIdList, mapIdArr...)
		}
		index := random.RandInt(0, len(mapIdList))
		mapId = mapIdList[index]
	}

	result := &data.MatchResult{
		MatchRoomId:    roomId,
		PlayId:         playId,
		MapId:          mapId,
		Members:        members,
		Robots:         robots,
		Players:        players,
		BattleRoomAddr: addr,
		IsFirstEnter:   isFirstEnter,
	}
	var rToken []byte
	if playId != 0 {
		rToken, err = m.authResult(result)
		if err != nil {
			m.logger.Error("matchmaking authResult err", zap.Any("error", err))
			if !isRetry {
				m.addRetryMatch(roomId, matchData)
			}
			return false
		}
	}
	matchResp := &pb.MatchResult{
		BattleRoomUrl: addr,
		BattleRoomId:  result.MatchRoomId,
		AuthByte:      rToken,
		PlayId:        playId,
		MapId:         mapId,
	}
	m.logger.Info("matchmaking success!", zap.Any("result", matchResp))
	//if matchByt, err := proto.Marshal(matchResp); err != nil {
	//	m.logger.Error("matchmaking MatchResponse marshal err", zap.Any("error", err))
	//} else {
	//	//resp := &notification.S2SNotifyMessage{
	//	//	MsgId: utils.NTF_MatchingSuccess,
	//	//	Data:  matchByt,
	//	//}
	//	//respByt, _ := proto.Marshal()
	//	option := miface.WithBytes(respByt)
	//	for _, datum := range matchData {
	//		for _, playerData := range datum.Members {
	//			topic := utils.MakeNotifyTopic(playerData.Uid)
	//			er := m.mq.Publish(topic, option)
	//			if er != nil {
	//				m.logger.Error("matchmaking publish err", zap.Any("error", er))
	//			}
	//		}
	//	}
	//	if mapData != nil {
	//		if mapData.MapId == mapId {
	//			mapData.MapCount += 1
	//		} else {
	//			mapData.MapId = mapId
	//			mapData.MapCount = 1
	//		}
	//		er := m.db.SaveBattleMap(playerId, playId, mapData)
	//		if er != nil {
	//			m.logger.Error("matchmaking SaveBattleMap error", zap.Error(er))
	//		}
	//	}
	//}
	return true
}

func (m *MatchManager) onMatchWithRivalSuccess(matchData []*model.MatchData) {
	members := make(map[string]int32)
	players := make(map[string]data.PlayerData)
	playId := matchData[0].PlayId
	uidsAllocate := make([]string, 0)
	for index, datum := range matchData {
		camp := int32(index + 1)
		for _, playerData := range datum.Members {
			members[playerData.Uid] = camp
			uidsAllocate = append(uidsAllocate, playerData.Uid)
			player := data.PlayerData{
				Uid:          playerData.Uid,
				Nickname:     playerData.Nickname,
				Avatar:       playerData.Avatar,
				HeroId:       playerData.HeroId,
				HeroLevel:    playerData.HeroLevel,
				SkinId:       playerData.SkinId,
				Attribute:    playerData.Attribute,
				PetProfileId: playerData.PetProfileId,
				PetSkill:     playerData.PetSkill,
			}
			players[player.Uid] = player
		}
	}
	roomId := uuid.New().String()
	addr, err := m.allocator.Allocate(playId, uidsAllocate...)
	if err != nil {
		m.logger.Error("matchmaking withRival allocateBattle err", zap.Any("error", err))
		return
	}
	result := &data.MatchResult{
		MatchRoomId:    roomId,
		PlayId:         playId,
		Members:        members,
		Robots:         make(map[int32][]int32),
		Players:        players,
		BattleRoomAddr: addr,
		IsFirstEnter:   true,
	}
	var rToken []byte
	if playId != 0 {
		rToken, err = m.authResult(result)
		if err != nil {
			m.logger.Error("matchmaking withRival authResult err", zap.Any("error", err))
			return
		}
	}
	matchResp := &pb.MatchResult{
		BattleRoomUrl: addr,
		BattleRoomId:  result.MatchRoomId,
		AuthByte:      rToken,
		PlayId:        playId,
	}
	m.logger.Info("matchmaking withRival success!", zap.Any("result", matchResp))
	//if matchByt, err := proto.Marshal(matchResp); err != nil {
	//	m.logger.Error("matchmaking withRival MatchResponse marshal err", zap.Any("error", err))
	//} else {
	//resp := &notification.S2SNotifyMessage{
	//	MsgId: utils.NTF_MatchingSuccess,
	//	Data:  matchByt,
	//}
	//respByt, _ := proto.Marshal(resp)
	//option := miface.WithBytes(respByt)
	//for _, datum := range matchData {
	//	for _, playerData := range datum.Members {
	//		topic := utils.MakeNotifyTopic(playerData.Uid)
	//		er := m.mq.Publish(topic, option)
	//		if er != nil {
	//			m.logger.Error("matchmaking withRival publish err", zap.Any("error", er))
	//		}
	//	}
	//}
	//}

}

func (m *MatchManager) authResult(result *data.MatchResult) ([]byte, error) {
	byt, err := json.Marshal(result)
	if err != nil {
		m.logger.Error("matchmaking MatchResponse marshal err", zap.Any("error", err))
		return nil, err
	}
	res := &pb2.AuthenticateRequest{
		Id:   result.MatchRoomId,
		Data: byt,
	}
	if authRes, err := m.aClient.Authenticate(context.TODO(), res); err != nil {
		m.logger.Error("matchmaking Authenticate err", zap.Any("error", err))
		return nil, err
	} else if bt, err := proto.Marshal(authRes); err != nil {
		m.logger.Error("matchmaking Authenticate marshal err", zap.Any("error", err))
		return nil, err
	} else {
		return bt, nil
	}
}

func (m *MatchManager) JoinMatch(groupSize int32, ticket []*pb.Ticket, playId int32, mapId []int32) {
	matchType := checkMatchType(groupSize)
	matchData := &model.MatchData{
		Members: make(map[string]*model.PlayerData),
	}
	var totalScore int32 = 0
	//TODO 暂时使用玩家ID作为队伍ID 暂时屏蔽浮动属性
	id := ticket[0].ProfileId
	for _, tk := range ticket {
		hero := &model.PlayerData{
			Uid:          tk.ProfileId,
			Nickname:     tk.Nickname,
			Avatar:       tk.Avatar,
			HeroId:       tk.DiffTag,
			HeroLevel:    tk.HeroLevel,
			HeroCups:     tk.HeroCups,
			SkinId:       tk.SkinId,
			Attribute:    tk.PetAttribute,
			Score:        tk.Score,
			PetProfileId: tk.PetProfileId,
			PetSkill:     tk.PetSkill,
			IsAgain:      tk.IsAgain,
		}
		matchData.Members[tk.ProfileId] = hero
		totalScore += tk.Score
	}
	matchData.Id = id
	matchData.GroupSize = groupSize
	matchData.Score = totalScore / int32(len(ticket))
	matchData.MatchTime = time.Now().Unix()
	matchData.PlayId = playId
	matchData.MapId = mapId

	if playId == 0 || playId == 999 {
		m.matchSingleRoom(matchData)
		return
	}
	playIdStr := strconv.Itoa(int(playId))
	if playId == 997 {
		matchType = string(db.Match_Type_Guid)
		m.matchFirstCombatRoom(matchData, playIdStr, matchType)
		return
	}
	m.playMap.LoadOrStore(playIdStr, 0)
	m.checkTeammate(matchData, matchType, 0)
}

func (m *MatchManager) JoinMatchWithRival(groupSize int32, ticket []*pb.Ticket, ticketRival []*pb.Ticket, playId int32) {
	matchData := &model.MatchData{
		Members: make(map[string]*model.PlayerData),
	}
	var totalScore int32 = 0
	//TODO 暂时使用玩家ID作为队伍ID 暂时屏蔽浮动属性
	id := ticket[0].ProfileId
	for _, tk := range ticket {
		hero := &model.PlayerData{
			Uid:       tk.ProfileId,
			Nickname:  tk.Nickname,
			Avatar:    tk.Avatar,
			HeroId:    tk.DiffTag,
			HeroLevel: tk.HeroLevel,
			SkinId:    tk.SkinId,
			//Attribute: tk.HeroAttribute,
			Score:        tk.Score,
			PetProfileId: tk.PetProfileId,
			PetSkill:     tk.PetSkill,
			IsAgain:      tk.IsAgain,
		}
		matchData.Members[tk.ProfileId] = hero
		totalScore += tk.Score
	}
	matchData.Id = id
	matchData.GroupSize = groupSize
	matchData.Score = totalScore / int32(len(ticket))
	matchData.MatchTime = time.Now().Unix()
	matchData.PlayId = playId
	//对手数据
	rivalMatchData := &model.MatchData{
		Members: make(map[string]*model.PlayerData),
	}
	rivalId := ticketRival[0].ProfileId
	for _, tk := range ticketRival {
		hero := &model.PlayerData{
			Uid:       tk.ProfileId,
			Nickname:  tk.Nickname,
			Avatar:    tk.Avatar,
			HeroId:    tk.DiffTag,
			HeroLevel: tk.HeroLevel,
			SkinId:    tk.SkinId,
			//Attribute: tk.HeroAttribute,
			Score:        tk.Score,
			PetProfileId: tk.PetProfileId,
			IsAgain:      tk.IsAgain,
		}
		rivalMatchData.Members[tk.ProfileId] = hero
		totalScore += tk.Score
	}
	rivalMatchData.Id = rivalId
	rivalMatchData.GroupSize = groupSize
	rivalMatchData.Score = totalScore / int32(len(ticket))
	rivalMatchData.MatchTime = time.Now().Unix()
	rivalMatchData.PlayId = playId
	m.onMatchWithRivalSuccess([]*model.MatchData{matchData, rivalMatchData})
}

func (m *MatchManager) JoinPVEMatch(groupSize int32, ticket []*pb.Ticket, playId int32, mapId []int32) {
	matchType := string(db.Match_Type_PVE)
	matchData := &model.MatchData{
		Members: make(map[string]*model.PlayerData),
	}
	var totalScore int32 = 0
	//TODO 暂时使用玩家ID作为队伍ID 暂时屏蔽浮动属性
	id := ticket[0].ProfileId
	for _, tk := range ticket {
		hero := &model.PlayerData{
			Uid:          tk.ProfileId,
			Nickname:     tk.Nickname,
			Avatar:       tk.Avatar,
			HeroId:       tk.DiffTag,
			HeroLevel:    tk.HeroLevel,
			HeroCups:     tk.HeroCups,
			SkinId:       tk.SkinId,
			Attribute:    tk.PetAttribute,
			Score:        tk.Score,
			PetProfileId: tk.PetProfileId,
			PetSkill:     tk.PetSkill,
			IsAgain:      tk.IsAgain,
		}
		matchData.Members[tk.ProfileId] = hero
		totalScore += tk.Score
	}
	matchData.Id = id
	matchData.GroupSize = groupSize
	matchData.Score = totalScore / int32(len(ticket))
	matchData.MatchTime = time.Now().Unix()
	matchData.PlayId = playId
	matchData.MapId = mapId
	playIdStr := strconv.Itoa(int(playId))
	m.pveMap.LoadOrStore(playIdStr, 0)
	m.checkPVETeammate(matchData, matchType, 0)
}

func (m *MatchManager) matchSingleRoom(data *model.MatchData) {
	m.onMatchSuccess([]*model.MatchData{data}, true)
}

func (m *MatchManager) matchFirstCombatRoom(data *model.MatchData, playId string, matchType string) {
	m.guidMap.LoadOrStore(playId, 0)
	//保存数据到新的队列
	err := m.db.PushData(matchType, string(db.Queue_Type_Wait), "0", playId, data)
	if err != nil {
		m.logger.Error("matchFirstCombatRoom error", zap.Error(err))
		return
	}
}

// CancelMatch 客户端主动发送取消匹配请求
func (m *MatchManager) CancelMatch(uid string) {
	keyParams, ok := m.playerDic.Load(uid)
	if !ok {
		m.logger.Error("cancel match fail,key not found", zap.String("uid", uid))
		return
	}
	params := keyParams.([]string)
	err := m.db.DeleteData(params[0], params[1], params[2], params[3], params[4])
	if err != nil {
		m.logger.Error("delete data", zap.Error(err))
		return
	}
}

func checkQueueLevel(subTime int64) int {
	if subTime <= utils.QUEUE_TIME_1 {
		return utils.QUEUE_LEVEL_0
	}
	if subTime <= utils.QUEUE_TIME_2 {
		return utils.QUEUE_LEVEL_3
	}
	if subTime <= utils.QUEUE_TIME_3 {
		return utils.QUEUE_LEVEL_3
	}
	if subTime <= utils.QUEUE_TIME_4 {
		return utils.QUEUE_LEVEL_3
	}
	return utils.QUEUE_LEVEL_3
}

func checkMatchType(groupSize int32) string {
	if groupSize == 1 {
		return string(db.Match_Type_Single)
	}
	if groupSize == 3 {
		return string(db.Match_Type_Team)
	}
	return string(db.Match_Type_Team)
}

func (m *MatchManager) CheckMatchStatus(uid string) int64 {
	var tim int64
	m.playMap.Range(func(key, value interface{}) bool {
		playId := key.(string)
		for i := 0; i < MaxLevel; i++ {
			queueLevel := strconv.Itoa(i)
			tim = m.matchStatus(uid, string(db.Match_Type_Team), queueLevel, playId)
			if tim > 0 {
				return false
			}
		}
		return true
	})
	if tim > 0 {
		return tim
	}
	m.pveMap.Range(func(key, value interface{}) bool {
		playId := key.(string)
		for i := 0; i < MaxLevel; i++ {
			queueLevel := strconv.Itoa(i)
			tim = m.matchStatus(uid, string(db.Match_Type_PVE), queueLevel, playId)
			if tim > 0 {
				return false
			}
		}
		return true
	})
	if tim > 0 {
		return tim
	}
	m.guidMap.Range(func(key, value interface{}) bool {
		playId := key.(string)
		for i := 0; i < MaxLevel; i++ {
			queueLevel := strconv.Itoa(i)
			tim = m.matchStatus(uid, string(db.Match_Type_Guid), queueLevel, playId)
			if tim > 0 {
				return false
			}
		}
		return true
	})
	return tim
}

func (m *MatchManager) matchStatus(uid string, matchType, queueLevel, playId string) int64 {
	dataMap, _ := m.db.GetAllWaitData(matchType, queueLevel, playId)
	if dataMap != nil {
		for _, matchData := range dataMap {
			if matchData.Members == nil {
				continue
			}
			if _, ok := matchData.Members[uid]; ok {
				return matchData.MatchTime
			}
		}
	}
	dataMap, _ = m.db.GetAllMatchData(matchType, queueLevel, playId)
	if dataMap != nil {
		for _, matchData := range dataMap {
			if matchData.Members == nil {
				continue
			}
			if _, ok := matchData.Members[uid]; ok {
				return matchData.MatchTime
			}
		}
	}
	return 0
}

func (m *MatchManager) updatePlayerDic(matchData *model.MatchData, matchType, queueType, queueLevel, playId string) {
	for uid := range matchData.Members {
		params := []string{matchType, queueType, queueLevel, playId, matchData.Id}
		m.playerDic.Store(uid, params)
	}

}

func (m *MatchManager) addRetryMatch(roomId string, match []*model.MatchData) {
	matchRetry := &model.MatchRetry{
		RoomId:        roomId,
		MatchData:     match,
		NextRetryTime: time.Now().UTC().Unix() + nextRetry[0],
	}
	m.retryMap.Store(roomId, matchRetry)
	for _, matchData := range match {
		for _, playerData := range matchData.Members {
			m.retryIdsMap.Store(playerData.Uid, roomId)
		}
	}
}

func (m *MatchManager) delRetryMatchByUid(uid string) {
	val, ok := m.retryIdsMap.Load(uid)
	if !ok {
		m.logger.Info("delRetryMatchByUid key not found", zap.String("uid", uid))
		return
	}
	roomId := val.(string)
	match, ok := m.retryMap.Load(roomId)
	if !ok {
		m.logger.Info("delRetryMatchByUid key not found", zap.String("roomId", roomId))
		return
	}
	m.retryMap.Delete(roomId)
	mh := match.(*model.MatchRetry)
	for _, datum := range mh.MatchData {
		for _, playerData := range datum.Members {
			m.retryIdsMap.Delete(playerData.Uid)
		}
	}
}

func (m *MatchManager) delRetryMatchByRoomId(roomId string) {
	match, ok := m.retryMap.Load(roomId)
	if !ok {
		m.logger.Info("delRetryMatchByRoomId key not found", zap.String("roomId", roomId))
		return
	}
	m.retryMap.Delete(roomId)
	mh := match.(*model.MatchRetry)
	for _, datum := range mh.MatchData {
		for _, playerData := range datum.Members {
			m.retryIdsMap.Delete(playerData.Uid)
		}
	}
}

func (m *MatchManager) pubCancelMatchMsg(uid []string) {
	for _, id := range uid {
		m.pushMqMsg(&pb.MatchingCancel{}, int32(utils.S2C_EVENT_S2C_MatchingCancel), id)
	}
}

func (m *MatchManager) pushMqMsg(res proto.Message, msgId int32, uid string) {
	//bt, _ := proto.Marshal(res)
	//notifyMsg := &notification.S2SNotifyMessage{
	//	MsgId: msgId,
	//	Data:  bt,
	//}
	//notifyData, _ := proto.Marshal(notifyMsg)
	//option := miface.WithBytes(notifyData)
	//notifyTopic := utils.MakeNotifyTopic(uid)
	//er := m.mq.Publish(notifyTopic, option)
	//if er != nil {
	//	m.logger.Error("mq publish err", zap.Any("error", er))
	//}
}

func (m *MatchManager) validRoom(roomId string) bool {
	battleRoomKey, err := db.MakeRoomKey(roomId)
	if err != nil {
		m.logger.Error("MatchManager validRoom NewKeyFromParts error", zap.Error(err))
		return false
	}
	resultCmd := m.db.Get(context.TODO(), battleRoomKey.String())
	if resultCmd.Err() != nil {
		m.logger.Error("RoomHeart redis Get error", zap.Error(resultCmd.Err()))
		return false
	}
	nowTime := time.Now().UTC().Unix()
	if heartTim, err := resultCmd.Int64(); err != nil {
		m.logger.Error("RoomHeart redis Val error", zap.Error(err))
		return false
	} else if nowTime-heartTim < 20 {
		return true
	}
	return false

}
