package internal

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"
)

type BattleLog struct {
	ID        string    `json:"id"`
	UID       string    `json:"uid"`
	PlayId    int32     `json:"play_id"`
	Cup       int32     `json:"cup"`
	Win       int32     `json:"win"`
	Mvp       int32     `json:"mvp"`
	HeroId    int32     `json:"hero_id"`
	SkinId    int32     `json:"skin_id"`
	BeginTime time.Time `json:"begin_time"`
	EndTime   time.Time `json:"end_time"`
}

func TestCkInsert(t *testing.T) {
	p := &Processor{}
	p.Init(zap.NewExample(), "./", "192.168.90.35:8900", "fr", "", "")
	err := p.insertData("battle_log", map[string]any{
		"id":         "123123123",
		"uid":        "dddddddddddd",
		"play_id":    123,
		"cup":        3,
		"win":        1,
		"mvp":        0,
		"hero_id":    10010,
		"skin_id":    1001001,
		"begin_time": time.Now().Add(-300 * time.Second),
		"end_time":   time.Now(),
	})
	fmt.Println("insert:", err)
}
