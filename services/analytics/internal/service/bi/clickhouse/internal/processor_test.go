package internal

import (
	"testing"
)

func TestProcessor_getType(t *testing.T) {
	tests := []struct {
		name string
		args any
		want string
	}{
		{
			name: "str",
			args: "1",
			want: "String",
		},
		{
			name: "int",
			args: 1,
			want: "Int32",
		},
		{
			name: "f64",
			args: 10.2,
			want: "Float64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{}
			if got := p.getType(tt.args); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}
