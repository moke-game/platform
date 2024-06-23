package test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/moke-game/platform.git/services/analytics/internal/service/bi"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi/mix_panel"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi/thinking_data"
)

type testCase struct {
	data  []byte
	valid bool
}

var testEventTypes = []bi.EventType{
	"Signed Up",
	bi.EventTypeUserSet,
	bi.EventTypeUserSetOnce,
	bi.EventTypeUserAdd,
	bi.EventTypeUserDel,
}

func iniTestEventProperties() []testCase {
	return []testCase{
		{
			data:  []byte(``),
			valid: true,
		},
		{
			data:  []byte(`[]`),
			valid: false,
		},
		{
			data:  []byte(`"":""`),
			valid: true,
			// "valid" but emptystring as an index is evil...
		},
		{
			data:  []byte(`"":`),
			valid: false,
		},
		{
			data:  []byte(`1:2`),
			valid: false,
		},
		{
			data:  []byte(`:""`),
			valid: false,
		},
		{
			data:  []byte(`{'a':'apple'}`),
			valid: false,
		},
		{
			data:  []byte(`[{"b":"banana"}]`),
			valid: false,
		},
		{
			data:  []byte(`"cherry"`),
			valid: false,
		},
		{
			data:  []byte(`date`),
			valid: false,
		},
		{
			data:  []byte(`"e":5.0,"f":false,"g":"grape","h":[1,2,3,4,5]`),
			valid: true,
		},
		{
			data:  []byte(`"i":123`),
			valid: true,
		},
		{
			data:  []byte(`"j":"juniper","k":"kumquat"`),
			valid: true,
		},
	}
}

func TestBiHandles(t *testing.T) {
	userID := "test"
	ip := "localhost"
	token := "test"
	testCase := iniTestEventProperties()
	if logger, err := zap.NewDevelopment(zap.AddStacktrace(zap.FatalLevel)); err != nil {
		t.Error(err)
	} else if queue, err := mock.NewLocalMessageQueue(logger, "development"); err != nil {
		t.Error(err)
	} else if dpTD, err := thinking_data.NewDataProcessor(logger, queue, userID, ip, token); err != nil {
		t.Error(err)
	} else if dpMP, err := mix_panel.NewDataProcessor(logger, queue, userID, ip, token); err != nil {
		t.Error(err)
	} else {
		for _, v := range testEventTypes {
			for k1, v1 := range testCase {
				err := dpTD.Handle(v, v1.data)
				if (err != nil) == v1.valid {
					t.Errorf("TestTD name:%s, case:#%d, Unexpected result '%t' for '%s'.\n error:%v",
						v,
						k1,
						(err != nil) == v1.valid,
						string(v1.data),
						err,
					)
				}

				err = dpMP.Handle(v, v1.data)
				if (err != nil) == v1.valid {
					t.Errorf("TestTD name:%s, case:#%d, Unexpected result '%t' for '%s'.\n error:%v",
						v,
						k1,
						(err != nil) == v1.valid,
						string(v1.data),
						err,
					)
				}
			}
		}
	}
}
