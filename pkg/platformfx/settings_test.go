package platformfx

import (
	"context"
	"testing"
	"time"

	"go.uber.org/fx"
)

type envResult struct {
	fx.Out

	Value string `name:"PlatformfxTestValue" envconfig:"PLATFORMFX_TEST_VALUE" default:"from-env"`
}

type envParams struct {
	fx.In

	Value string `name:"PlatformfxTestValue"`
}

func TestProvideFromEnv(t *testing.T) {
	t.Setenv("PLATFORMFX_TEST_VALUE", "ok")

	var got string
	app := fx.New(
		fx.NopLogger,
		ProvideFromEnv[envResult](),
		fx.Invoke(func(p envParams) { got = p.Value }),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Err(); err != nil {
		t.Fatalf("fx graph: %v", err)
	}
	if err := app.Start(ctx); err != nil {
		t.Fatalf("start: %v", err)
	}
	if err := app.Stop(ctx); err != nil {
		t.Fatalf("stop: %v", err)
	}
	if got != "ok" {
		t.Fatalf("Value = %q, want ok", got)
	}
}
