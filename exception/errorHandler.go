package exception

import (
	"bara-playdate-api/constant"
	"bara-playdate-api/model/result"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, isValidation := err.(ValidationError)
	if isValidation {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return ctx.JSON(result.ResponseResult{
			ResponseCode:        http.StatusBadRequest,
			ResponseDescription: constant.BAD_REQUEST,
			ResponseTime:        time.Now().Format("2006-01-02 15:04:05"),
			ResponseDatas:       messages,
		})
	}

	_, isNotFound := err.(NotFoundError)
	if isNotFound {
		return ctx.JSON(result.ResponseResult{
			ResponseCode:        http.StatusNotFound,
			ResponseDescription: constant.NOT_FOUND,
			ResponseTime:        time.Now().Format("2006-01-02 15:04:05"),
			ResponseDatas:       err.Error(),
		})
	}

	_, isUnauthorized := err.(UnauthorizedError)
	if isUnauthorized {
		return ctx.JSON(result.ResponseResult{
			ResponseCode:        http.StatusUnauthorized,
			ResponseDescription: constant.UNAUTHORIZED,
			ResponseTime:        time.Now().Format("2006-01-02 15:04:05"),
			ResponseDatas:       err.Error(),
		})
	}

	return ctx.JSON(result.ResponseResult{
		ResponseCode:        http.StatusInternalServerError,
		ResponseDescription: constant.GENERAL_ERROR,
		ResponseTime:        time.Now().Format("2006-01-02 15:04:05"),
		ResponseDatas:       err.Error(),
	})
}
