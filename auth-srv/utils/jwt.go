package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// BizType ...
type BizType int

const (
	// BizStruct ...
	BizStruct BizType = iota
	// BizMap ...
	BizMap
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

// NewMapToken generate tokens used for auth
func (j *Jwt) NewMapToken(params map[string]interface{}, expire time.Time) (string, error) {
	var (
		claims    jwt.Claims
		jwtSecret = j.ByteSecret()
	)
	nowTime := time.Now()
	if expire.Before(nowTime) {
		expire = nowTime.Add(3 * time.Hour)
	}

	mapClaims := make(jwt.MapClaims)
	for k, v := range params {
		mapClaims[k] = v
	}
	mapClaims["exp"] = expire.Unix()
	mapClaims["iss"] = j.Issuer
	claims = mapClaims

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseTokenMap parsing token
func (j *Jwt) ParseTokenMap(token string) (claims jwt.Claims, err error) {
	var (
		tokenClaims *jwt.Token
		jwtSecret   = j.ByteSecret()
	)
	claims = make(jwt.MapClaims)

	tokenClaims, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
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
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// Parse parsing token
func (j *Jwt) Parse(tokenString string) (claims jwt.Claims, err error) {
	var (
		token     *jwt.Token
		jwtSecret = j.ByteSecret()
	)

	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims = mapClaims
	} else if myClaims, ok := token.Claims.(MyClaims); ok && token.Valid {
		claims = myClaims
	}

	return claims, nil
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
