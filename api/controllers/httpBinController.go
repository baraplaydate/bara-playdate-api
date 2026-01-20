package controller

import (
	"net/http"

	"bara-playdate-api/constant"
	"bara-playdate-api/model"
	"bara-playdate-api/model/result"
	"bara-playdate-api/resClient"
	"bara-playdate-api/utils"

	"github.com/gofiber/fiber/v2"
)

type HttpBinController struct {
}

func NewHttpBinController() *HttpBinController {
	return &HttpBinController{}
}

func (controller HttpBinController) Route(app *fiber.App) {
	app.Get("/v1/api/httpbin", controller.PostHttpBin)
}

func (controller HttpBinController) PostHttpBin(c *fiber.Ctx) error {

	httpbinRestClient := resClient.NewHttpBinRestClient()
	httpBin := model.HttpBin{
		Name: "bayuwidiasantoso",
	}
	var response map[string]interface{}
	httpbinRestClient.PostMethod(c.Context(), &httpBin, &response)
	utils.NewLogger().Info("log response service ", response)

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       nil,
	})
}
