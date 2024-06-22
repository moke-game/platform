package model

// MatchRetry 匹配队列
type MatchRetry struct {
	RoomId        string
	MatchData     []*MatchData //匹配队列 重试
	NextRetryTime int64        //下一次重试时间戳
	RetryCount    int          // 重试次数
}

func NewMatchRetry() *MatchRetry {
	matchRetry := &MatchRetry{
		MatchData: make([]*MatchData, 0),
	}
	return matchRetry
}

func (m *MatchRetry) AddRetryCount() {
	m.RetryCount++
}

func (m *MatchRetry) SetNextTime(tim int64) {
	m.NextRetryTime = tim
}
