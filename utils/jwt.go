package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateToken(secret string, data Payload) (string, error) {
	header := header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerJsonByte, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadJsonByte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	headerB64 := encodeBase64String(headerJsonByte)
	payloadB64 := encodeBase64String(payloadJsonByte)

	headerAndPayloadHashString := headerB64 + "." + payloadB64

	byteSecret := []byte(secret)
	byteHeaderAndPayload := []byte(headerAndPayloadHashString)

	hash := hmac.New(sha256.New, byteSecret)
	hash.Write(byteHeaderAndPayload)

	signature := hash.Sum(nil)

	signatureB64 := encodeBase64String(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil

}

func encodeBase64String(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

func Verify(jwt, secret string) (bool, error) {
	jwtChunkArr := strings.Split(jwt, ".")
	headerAndPayloadHash := jwtChunkArr[0] + "." + jwtChunkArr[1]

	byteSecret := []byte(secret)
	byteHeaderAndPayload := []byte(headerAndPayloadHash)

	hash := hmac.New(sha256.New, byteSecret)
	hash.Write(byteHeaderAndPayload)

	signature := hash.Sum(nil)

	signatureB64 := encodeBase64String(signature)

	verifyJwt := headerAndPayloadHash + "." + signatureB64

	if jwt == verifyJwt {
		return true, nil
	}
	return false, errors.New("invalid token")
}

func DecodeToken(token, secret string) (*Payload, error) {

	verified, err := Verify(token, secret)
	if err != nil {
		return nil, err
	}
	if !verified {
		return nil, errors.New("invalid token")
	}
	jwtChunkArr := strings.Split(token, ".")
	// headerHash := jwtChunkArr[0]
	payloadHash := jwtChunkArr[1]

	payloadByte, err := base64.RawURLEncoding.DecodeString(payloadHash)
	if err != nil {
		return nil, err
	}

	var payload Payload
	err = json.Unmarshal(payloadByte, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
