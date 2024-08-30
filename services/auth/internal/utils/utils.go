package utils

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
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
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenSigningMethod
		}
		return []byte(key), nil
	})
	var ve *jwt.ValidationError
	if errors.As(err, &ve) {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", nil, ErrTokenMalformed
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return "", nil, ErrTokenExpired
		} else {
			return "", nil, ErrTokenHandle
		}
	}

	if claims, ok := claim.Claims.(jwt.MapClaims); ok && claim.Valid {
		if tp, ok := claims["type"]; ok && tp.(float64) == float64(tokenType) {
			resUid := ""
			var resData []byte
			if uid, ok := claims["uid"].(string); ok {
				resUid = uid
			}
			if data, ok := claims["data"].(string); ok {
				// base64 decode
				if raw, err := base64.StdEncoding.DecodeString(data); err != nil {
					return "", nil, ErrTokenHandle
				} else {
					resData = raw
				}
			}
			return resUid, resData, nil
		}
	}
	return "", nil, ErrTokenHandle
}
