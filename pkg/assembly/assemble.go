// Package assembly starts platform processes from short names.
//
//	assembly.Main("profile")
//	assembly.Main("profile", "knapsack", "chat")
//	assembly.Main("platform")
//	assembly.Main("认证", "玩家", "背包")
//
// A game that only dials platform:
//
//	assembly.Main("auth.mw", "profile.client", "knapsack.client")
//
// Or list typed bricks:
//
//	fxmain.Main(assembly.NATS, assembly.Cache, assembly.AuthMW, assembly.Profile)
package assembly

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/gstones/moke-kit/fxmain"
	"go.uber.org/fx"
)

// Main expands names and runs fxmain.Main. Unknown names exit the process.
func Main(names ...string) {
	opts, err := Options(names...)
	if err != nil {
		log.Fatal(err)
	}
	fxmain.Main(opts...)
}

// Must expands names or panics. Use next to extra fx options:
//
//	fxmain.Main(assembly.Must("profile"), extra)
func Must(names ...string) fx.Option {
	opts, err := Options(names...)
	if err != nil {
		panic(err)
	}
	return fx.Options(opts...)
}

// Options expands names to fx options (deduped; auth resolved to one module).
func Options(names ...string) ([]fx.Option, error) {
	resolved, err := Resolve(names...)
	if err != nil {
		return nil, err
	}
	opts := make([]fx.Option, 0, len(resolved))
	for _, name := range resolved {
		opts = append(opts, bricks[name])
	}
	return opts, nil
}

// Resolve returns the ordered brick names after recipe expand, alias, dedupe, auth pick.
func Resolve(names ...string) ([]string, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("assembly: need at least one name (have: %s)", recipeList())
	}
	seen := make(map[string]struct{}, 16)
	ordered := make([]string, 0, 16)
	for _, raw := range names {
		parts, err := expand(raw)
		if err != nil {
			return nil, err
		}
		for _, part := range parts {
			if _, ok := seen[part]; ok {
				continue
			}
			seen[part] = struct{}{}
			ordered = append(ordered, part)
		}
	}
	return resolveAuth(ordered), nil
}

// Names lists recipe names (the usual Main() arguments).
func Names() []string {
	out := make([]string, 0, len(recipes))
	for name := range recipes {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

func expand(raw string) ([]string, error) {
	name := canon(raw)
	if parts, ok := recipes[name]; ok {
		return parts, nil
	}
	if _, ok := bricks[name]; ok {
		return []string{name}, nil
	}
	return nil, fmt.Errorf("assembly: unknown %q (recipes: %s)", raw, recipeList())
}

func canon(name string) string {
	n := strings.ToLower(strings.TrimSpace(name))
	if alias, ok := aliases[n]; ok {
		return alias
	}
	return n
}

func resolveAuth(names []string) []string {
	hasAuth, hasMW, hasPriv := false, false, false
	out := make([]string, 0, len(names))
	for _, name := range names {
		switch name {
		case "auth":
			hasAuth = true
		case "auth.mw":
			hasMW = true
		case "auth.private":
			hasPriv = true
		default:
			out = append(out, name)
		}
	}
	switch {
	case hasAuth:
		out = append(out, "auth")
	case hasMW:
		out = append(out, "auth.mw")
	case hasPriv:
		out = append(out, "auth.private")
	}
	return out
}

func recipeList() string {
	return strings.Join(Names(), ", ")
}
