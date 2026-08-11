package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestCreatAndParseJwt(t *testing.T) {
	t.Parallel()
	const key = "test-secret"
	tok, err := CreatJwt("uid-1", TokenTypeAccess, key, []byte("payload"), time.Hour)
	if err != nil {
		t.Fatalf("CreatJwt: %v", err)
	}
	uid, data, err := ParseToken(tok, TokenTypeAccess, key)
	if err != nil {
		t.Fatalf("ParseToken: %v", err)
	}
	if uid != "uid-1" {
		t.Fatalf("uid=%q", uid)
	}
	if string(data) != "payload" {
		t.Fatalf("data=%q", data)
	}
}

func TestParseTokenExpired(t *testing.T) {
	t.Parallel()
	const key = "test-secret"
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  "uid-1",
		"type": float64(TokenTypeAccess),
		"exp":  time.Now().Add(-time.Hour).Unix(),
	})
	tok, err := at.SignedString([]byte(key))
	if err != nil {
		t.Fatalf("SignedString: %v", err)
	}
	_, _, err = ParseToken(tok, TokenTypeAccess, key)
	if err != ErrTokenExpired {
		t.Fatalf("err=%v want ErrTokenExpired", err)
	}
}

func TestParseTokenWrongType(t *testing.T) {
	t.Parallel()
	const key = "test-secret"
	tok, err := CreatJwt("uid-1", TokenTypeAccess, key, nil, time.Hour)
	if err != nil {
		t.Fatalf("CreatJwt: %v", err)
	}
	_, _, err = ParseToken(tok, TokenTypeRefresh, key)
	if err != ErrTokenHandle {
		t.Fatalf("err=%v want ErrTokenHandle", err)
	}
}
