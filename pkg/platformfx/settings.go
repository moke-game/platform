package platformfx

import (
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
)

// ProvideFromEnv loads T from environment tags and provides it to the fx graph.
// T should be an fx.Out settings result (envconfig + name tags).
func ProvideFromEnv[T any]() fx.Option {
	return fx.Provide(func() (out T, err error) {
		err = utility.Load(&out)
		return
	})
}
