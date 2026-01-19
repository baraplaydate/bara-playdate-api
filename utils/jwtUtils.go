package utils

import (
	"bara-playdate-api/exception"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateToken(username string, roles []map[string]interface{}, config Config) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(3 * 24 * time.Hour).Unix(),
		"iat":        time.Now().Unix(),
		"iss":        config.ApiKeyEncode,
		"authorized": true,
		"jti":        uuid.New().String(),
	})

	signatureKey, _ := base64.StdEncoding.DecodeString(config.SignatureKeyEncode)
	apiKey, _ := base64.StdEncoding.DecodeString(config.ApiKeyEncode)

	hash := HmacEncode(signatureKey, []byte(apiKey))

	tokenSigned, err := token.SignedString([]byte(hash))
	exception.PanicLogging(err)

	return tokenSigned
}
