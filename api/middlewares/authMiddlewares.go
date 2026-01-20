package middleware

import (
	"net/http"
	"strconv"

	"bara-playdate-api/constant"
	"bara-playdate-api/model/result"
	"bara-playdate-api/utils"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthenticationJWT(config utils.Config) func(*fiber.Ctx) error {

	signatureKey := []byte(config.SignatureKey)
	apiKey := []byte(config.ApiKey)

	hash := utils.HmacEncode(signatureKey, apiKey)

	return jwtware.New(jwtware.Config{

		SigningKey: []byte(hash),
		SuccessHandler: func(ctx *fiber.Ctx) error {

			headerApiKey := ctx.Get("Api-Key")
			headerSignature := ctx.Get("Signature")
			headerSignatureTime := ctx.Get("Signature-Time")

			if headerApiKey != config.ApiKeyEncode {
				return ctx.Status(fiber.StatusUnauthorized).
					JSON(result.ResponseResult{
						ResponseCode:        http.StatusUnauthorized,
						ResponseDescription: constant.UNAUTHORIZED,
						ResponseTime:        utils.DateToStdNow(),
						ResponseDatas:       "Missing or invalid Api-Key header",
					})
			}

			if headerSignature != config.SignatureKeyEncode {
				return ctx.Status(fiber.StatusUnauthorized).
					JSON(result.ResponseResult{
						ResponseCode:        http.StatusUnauthorized,
						ResponseDescription: constant.UNAUTHORIZED,
						ResponseTime:        utils.DateToStdNow(),
						ResponseDatas:       "Missing or invalid Signature header",
					})
			}

			signatureTime, err := strconv.Atoi(headerSignatureTime)
			if err != nil || signatureTime < 0 {
				return ctx.Status(fiber.StatusUnauthorized).
					JSON(result.ResponseResult{
						ResponseCode:        http.StatusUnauthorized,
						ResponseDescription: constant.UNAUTHORIZED,
						ResponseTime:        utils.DateToStdNow(),
						ResponseDatas:       "Missing or invalid Signature-Time header",
					})
			}

			if ok, err := utils.HmacDecode(signatureKey, apiKey, hash); !ok {
				if err != nil {
					return ctx.Status(fiber.StatusUnauthorized).
						JSON(result.ResponseResult{
							ResponseCode:        http.StatusUnauthorized,
							ResponseDescription: constant.UNAUTHORIZED,
							ResponseTime:        utils.DateToStdNow(),
							ResponseDatas:       "Missing or invalid Signature-Time header",
						})
				}
			}

			return ctx.Next()

		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return ctx.Status(fiber.StatusBadRequest).
					JSON(result.ResponseResult{
						ResponseCode:        http.StatusBadRequest,
						ResponseDescription: constant.BAD_REQUEST,
						ResponseTime:        utils.DateToStdNow(),
						ResponseDatas:       "Missing or malformed JWT",
					})
			} else {
				return ctx.Status(fiber.StatusUnauthorized).
					JSON(result.ResponseResult{
						ResponseCode:        http.StatusUnauthorized,
						ResponseDescription: constant.UNAUTHORIZED,
						ResponseTime:        utils.DateToStdNow(),
						ResponseDatas:       "Invalid or expired JWT",
					})

			}
		},
	})
}
