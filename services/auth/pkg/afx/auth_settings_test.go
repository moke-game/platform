package afx

import "testing"

func TestValidateJwtExpire(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		deployment string
		expire     int32
		wantErr    bool
	}{
		{name: "prod positive", deployment: "prod", expire: 12},
		{name: "prod zero", deployment: "prod", expire: 0, wantErr: true},
		{name: "prod negative", deployment: "prod_gcp", expire: -1, wantErr: true},
		{name: "local zero ok", deployment: "local", expire: 0},
		{name: "dev zero ok", deployment: "dev", expire: 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateJwtExpire(tc.deployment, tc.expire)
			if tc.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
