package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWTClaims(sub, iss, aud string, exp time.Duration) jwt.Claims {
	return jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour * 24 * exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": iss,
		"aud": aud,
	}
}

func GenerateToken(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token, aud, iss, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(aud),
		jwt.WithIssuer(iss),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
