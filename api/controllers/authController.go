package controller

import (
	"bara-playdate-api/exception"
	"bara-playdate-api/repository"
	"bara-playdate-api/utils"
	"net/http"
	"strings"

	"bara-playdate-api/constant"
	"bara-playdate-api/model/criteria"
	"bara-playdate-api/model/result"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthController(authRepository *repository.AuthRepository, config utils.Config) *AuthController {
	return &AuthController{AuthRepository: *authRepository, Config: config}
}

type AuthController struct {
	repository.AuthRepository
	utils.Config
}

func (controller AuthController) Route(app *fiber.App) {

	group := app.Group(controller.Config.Route)

	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, welcome to Web Service " + app.Config().AppName + "!")
	})

	group.Post("/login", controller.Authentication)

}

func (controller AuthController) Authentication(c *fiber.Ctx) error {
	var authCriteria criteria.AuthCriteria
	err := c.BodyParser(&authCriteria)
	exception.PanicLogging(err)

	usernameOrEmail := strings.TrimSpace(*authCriteria.Username)
	password := strings.TrimSpace(*authCriteria.Password)

	userResult, err := controller.AuthRepository.Authentication(c.Context(), usernameOrEmail)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}

	// proses encrypt password ini di disable sementara karena penyesuaian dengan frontend
	// passwordDecrypt, errEncrypt := utils.DecryptAes256Sha256([]byte(*authCriteria.Password), constant.KEY_PASS_AES)
	// if errEncrypt != nil {
	// 	utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", errEncrypt.Error())
	// 	panic(exception.UnauthorizedError{
	// 		Message: errEncrypt.Error(),
	// 	})
	// }

	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(password))
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", "incorrect username and password")
		panic(exception.UnauthorizedError{
			Message: "incorrect username and password",
		})
	}

	tokenJwtResult := utils.GenerateToken(userResult.Username, nil, controller.Config)

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       tokenJwtResult,
	})
}
