package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Jwt ...
type Jwt struct {
	Secret  string
	Issuer  string
	Expires time.Duration
}

// JwtToken ...
type JwtToken struct {
	Token     string
	ExpiresAt time.Time
}

// ByteSecret ...
func (j *Jwt) ByteSecret() []byte {
	return []byte(j.Secret)
}

// MyClaims ...
type MyClaims struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// NewToken generate tokens used for auth
func (j *Jwt) NewToken(claims MyClaims) (JwtToken, error) {
	var (
		jwtToken  JwtToken
		jwtSecret = j.ByteSecret()
		nowTime   = time.Now()
		expireAt  time.Time
	)

	if j.Expires <= 0 {
		j.Expires = 3 * time.Hour
	}

	expireAt = nowTime.Add(j.Expires)
	jwtToken.ExpiresAt = expireAt

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireAt.Unix(),
		Issuer:    j.Issuer,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	jwtToken.Token = token
	return jwtToken, err
}

// ParseToken parsing token
func (j *Jwt) ParseToken(token string) (*MyClaims, error) {
	var (
		jwtSecret = j.ByteSecret()
	)
	
	// Parse the token
	tokenType, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*MyClaims); ok && tokenType.Valid {
		return claims, nil
	}
	return nil, err
}
