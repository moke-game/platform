package assembly

import (
	"slices"
	"strings"
	"testing"
)

func TestResolveRecipes(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in   []string
		want []string
	}{
		{[]string{"profile"}, []string{"nats", "cache", "profile", "auth.mw"}},
		{[]string{"auth"}, []string{"cache", "auth"}},
		{[]string{"analytics"}, []string{"analytics", "auth.private"}},
		{[]string{"room"}, []string{"agones", "room"}},
		{[]string{"platform"}, []string{
			"nats", "cache", "analytics", "profile", "knapsack", "mail",
			"party", "buddy", "leaderboard", "chat", "matchmaking", "auth",
		}},
	}
	for _, tc := range cases {
		got, err := Resolve(tc.in...)
		if err != nil {
			t.Fatalf("Resolve(%v): %v", tc.in, err)
		}
		if !slices.Equal(got, tc.want) {
			t.Fatalf("Resolve(%v) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestResolveComposeAndAuth(t *testing.T) {
	t.Parallel()
	got, err := Resolve("profile", "knapsack", "chat")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"nats", "cache", "profile", "knapsack", "chat", "auth.mw"}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	got, err = Resolve("auth", "profile")
	if err != nil {
		t.Fatal(err)
	}
	if slices.Contains(got, "auth.mw") || slices.Contains(got, "auth.private") {
		t.Fatalf("auth should replace mw/private: %v", got)
	}
	if !slices.Contains(got, "auth") || !slices.Contains(got, "profile") || !slices.Contains(got, "nats") {
		t.Fatalf("missing bricks: %v", got)
	}

	got, err = Resolve("analytics", "profile")
	if err != nil {
		t.Fatal(err)
	}
	if slices.Contains(got, "auth.private") {
		t.Fatalf("public+analytics must not keep pass-through: %v", got)
	}
	if !slices.Contains(got, "auth.mw") {
		t.Fatalf("public+analytics should keep auth.mw: %v", got)
	}
}

func TestResolveAliases(t *testing.T) {
	t.Parallel()
	got, err := Resolve("认证", "玩家", "背包")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"auth", "profile", "knapsack", "nats", "cache"} {
		if !slices.Contains(got, name) {
			t.Fatalf("alias assemble missing %s: %v", name, got)
		}
	}
	if slices.Contains(got, "auth.mw") {
		t.Fatalf("认证 expands to auth, which replaces auth.mw: %v", got)
	}

	got, err = Resolve("JWT", "Profile")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(got, "auth.mw") || !slices.Contains(got, "profile") {
		t.Fatalf("case-insensitive alias: %v", got)
	}
}

func TestResolveClientBrick(t *testing.T) {
	t.Parallel()
	got, err := Resolve("auth.mw", "profile.client", "knapsack.client")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"profile.client", "knapsack.client", "auth.mw"}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestResolveErrors(t *testing.T) {
	t.Parallel()
	if _, err := Resolve(); err == nil {
		t.Fatal("empty names should error")
	}
	_, err := Resolve("no-such-brick")
	if err == nil || !strings.Contains(err.Error(), "unknown") {
		t.Fatalf("unknown brick: %v", err)
	}
}

func TestMustAndOptions(t *testing.T) {
	t.Parallel()
	opts, err := Options("profile")
	if err != nil {
		t.Fatal(err)
	}
	if len(opts) != 4 {
		t.Fatalf("Options(profile) len = %d, want 4", len(opts))
	}
	if Must("mail") == nil {
		t.Fatal("Must returned nil")
	}
}

func TestNames(t *testing.T) {
	t.Parallel()
	got := Names()
	if !slices.Contains(got, "platform") || !slices.Contains(got, "profile") {
		t.Fatalf("Names() = %v", got)
	}
	if !slices.IsSorted(got) {
		t.Fatalf("Names not sorted: %v", got)
	}
}

func TestRecipeBricksExist(t *testing.T) {
	t.Parallel()
	for name, parts := range recipes {
		for _, part := range parts {
			if _, ok := bricks[part]; !ok {
				t.Fatalf("recipe %q lists unknown brick %q", name, part)
			}
		}
	}
	for alias, target := range aliases {
		if _, ok := recipes[target]; ok {
			continue
		}
		if _, ok := bricks[target]; ok {
			continue
		}
		t.Fatalf("alias %q -> unknown %q", alias, target)
	}
}
