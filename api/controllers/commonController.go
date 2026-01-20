package controller

import (
	"encoding/json"
	"net/http"

	"bara-playdate-api/constant"
	"bara-playdate-api/exception"
	"bara-playdate-api/model/result"
	"bara-playdate-api/utils"

	"github.com/gofiber/fiber/v2"
)

type CommonController struct {
	utils.Config
}

func NewCommonController(config utils.Config) *CommonController {
	return &CommonController{Config: config}
}

func (controller CommonController) Route(app *fiber.App) {

	group := app.Group(controller.Config.Route)

	group.Post("/encrypt", controller.GetEncryptAES)
	group.Post("/decrypt", controller.GetDecryptAES)
	group.Post("/encryptIsJson", controller.GetEncryptAESIsJson)
	group.Post("/decryptIsJson", controller.GetDecryptAESIsJson)
}

func (controller CommonController) GetEncryptAES(c *fiber.Ctx) error {
	typeKey := c.FormValue("typeKey")
	password := ""

	if typeKey == "DATA" {
		password = constant.KEY_AES
	} else if typeKey == "PASSWORD" {
		password = constant.KEY_PASS_AES
	}

	resultEncrypt, err := utils.EncryptAes256Sha256(c.FormValue("param"), password)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       resultEncrypt,
	})
}

func (controller CommonController) GetDecryptAES(c *fiber.Ctx) error {
	typeKey := c.FormValue("typeKey")
	password := ""

	if typeKey == "DATA" {
		password = constant.KEY_AES
	} else if typeKey == "PASSWORD" {
		password = constant.KEY_PASS_AES
	}

	resultDecrypt, err := utils.DecryptAes256Sha256([]byte(c.FormValue("param")), password)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       resultDecrypt,
	})
}

func (controller CommonController) GetEncryptAESIsJson(c *fiber.Ctx) error {

	var jsonData map[string]interface{}

	err := c.BodyParser(&jsonData)
	exception.PanicLogging(err)

	jsonDataBytes, err := json.Marshal(jsonData["data"])
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	resultEncrypt, errEncrypt := utils.EncryptAes256Sha256(string(jsonDataBytes), constant.KEY_AES)
	if errEncrypt != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", errEncrypt.Error())
		panic(exception.ValidationError{
			Message: errEncrypt.Error(),
		})
	}

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       resultEncrypt,
	})
}

func (controller CommonController) GetDecryptAESIsJson(c *fiber.Ctx) error {
	var jsonRequest map[string]interface{}

	err := c.BodyParser(&jsonRequest)
	exception.PanicLogging(err)

	encryptedData, err := json.Marshal(jsonRequest["encryptedData"])
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	resultDecrypt, errEncrypt := utils.DecryptAes256Sha256(encryptedData, constant.KEY_AES)
	if errEncrypt != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", errEncrypt.Error())
		panic(exception.ValidationError{
			Message: errEncrypt.Error(),
		})
	}

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       resultDecrypt,
	})
}
