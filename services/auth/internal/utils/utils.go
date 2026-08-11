package utils

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType int32

const (
	TokenTypeAccess TokenType = iota
	TokenTypeRefresh
)

var (
	ErrTokenSigningMethod = errors.New("ErrTokenSigningMethod")
	ErrTokenExpired       = errors.New("ErrTokenExpired")
	ErrTokenMalformed     = errors.New("ErrTokenMalformed")
	ErrTokenHandle        = errors.New("ErrTokenHandle")
	ErrSignedString       = errors.New("ErrSignedString")
)

// CreatJwt 生成一个JwtToken，包含uid
func CreatJwt(uid string, tp TokenType, key string, data []byte, duration time.Duration) (string, error) {
	exp := int64(0)
	if duration > 0 {
		exp = time.Now().Add(duration).Unix()
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"type": tp,
		"exp":  exp,
		"data": data,
	})
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", ErrSignedString
	}
	return token, nil
}

// ParseToken 从Jwt中解析Token
func ParseToken(token string, tokenType TokenType, key string) (string, []byte, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenSigningMethod
		}
		return []byte(key), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return "", nil, ErrTokenMalformed
		case errors.Is(err, jwt.ErrTokenExpired), errors.Is(err, jwt.ErrTokenNotValidYet):
			return "", nil, ErrTokenExpired
		default:
			return "", nil, ErrTokenHandle
		}
	}

	claims, ok := claim.Claims.(jwt.MapClaims)
	if !ok || !claim.Valid {
		return "", nil, ErrTokenHandle
	}
	tp, ok := claims["type"]
	if !ok || tp.(float64) != float64(tokenType) {
		return "", nil, ErrTokenHandle
	}

	resUid := ""
	var resData []byte
	if uid, ok := claims["uid"].(string); ok {
		resUid = uid
	}
	if data, ok := claims["data"].(string); ok {
		raw, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return "", nil, ErrTokenHandle
		}
		resData = raw
	}
	return resUid, resData, nil
}
