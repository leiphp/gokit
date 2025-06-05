package jwt

import (
	"time"
)

var Secret = []byte("your-secret")

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func Generate(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(Secret)
}

func Parse(tokenStr string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
