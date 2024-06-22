package model

// 匹配队列
type MatchQueue struct {
	matchQueue map[string]*MatchData //匹配队列 满员队伍
	waitQueue  map[string]*MatchData //等待队列 不满员的队伍进入等待队列
}

func NewMatchQueue() *MatchQueue {
	matchQueue := &MatchQueue{
		matchQueue: make(map[string]*MatchData),
		waitQueue:  make(map[string]*MatchData),
	}
	return matchQueue
}

func (m *MatchQueue) addMatch(match *MatchData) {
	m.matchQueue[match.Id] = match
}

func (m *MatchQueue) removeMatch(id string) {
	delete(m.matchQueue, id)
}

func (m *MatchQueue) addWait(match *MatchData) {
	m.waitQueue[match.Id] = match
}

func (m *MatchQueue) removeWait(id string) {
	delete(m.waitQueue, id)
}
