package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/akramboussanni/myrae/config"
)

type Credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Jwt struct {
	Header  JwtHeader
	Payload Claims
}

func (jwt Jwt) generateToken() string {
	header, _ := json.Marshal(jwt.Header)
	payload, _ := json.Marshal(jwt.Payload)

	data := base64.RawURLEncoding.EncodeToString(header) + "." + base64.RawURLEncoding.EncodeToString(payload)

	h := hmac.New(sha256.New, config.JwtSecret)
	h.Write([]byte(data))
	rawSig := h.Sum(nil)

	return data + "." + base64.RawURLEncoding.EncodeToString(rawSig)
}

type JwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type Claims struct {
	Id         string `json:"sub"`
	IssuedAt   int64  `json:"iat"`
	Expiration int64  `json:"exp"`
	Email      string `json:"email"`
	Role       string `json:"role"`
}
