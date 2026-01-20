package controller

import (
	"net/http"
	"strings"

	middleware "bara-playdate-api/api/middlewares"
	"bara-playdate-api/constant"
	"bara-playdate-api/exception"
	"bara-playdate-api/model/criteria"
	"bara-playdate-api/model/entity"
	"bara-playdate-api/model/result"
	"bara-playdate-api/repository"
	"bara-playdate-api/utils"
	"bara-playdate-api/utils/paginate"
	"bara-playdate-api/validation"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	repository.UserRepository
	utils.Config
}

func NewUserController(userRepository *repository.UserRepository, config utils.Config) *UserController {
	return &UserController{UserRepository: *userRepository, Config: config}
}

func (controller UserController) Route(app *fiber.App) {

	group := app.Group(controller.Config.Route)

	group.Post("/user/store", middleware.AuthenticationJWT(controller.Config), controller.Create)
	group.Put("/user/update/:id", middleware.AuthenticationJWT(controller.Config), controller.Update)
	group.Put("/user/updateIsActive/:id", middleware.AuthenticationJWT(controller.Config), controller.UpdateIsActive)
	group.Post("/user/getAllDataUserByParam", middleware.AuthenticationJWT(controller.Config), controller.GetAllDataUserByParam)
	group.Get("/user/getDataById/:id", middleware.AuthenticationJWT(controller.Config), controller.FindById)
	group.Get("/user", middleware.AuthenticationJWT(controller.Config), controller.FindAll)
}

func (controller UserController) Create(c *fiber.Ctx) error {
	var requestModel criteria.StoreUserCriteria

	err := c.BodyParser(&requestModel)
	exception.PanicLogging(err)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
	}

	validation.ValidateCriteria(requestModel)

	genPass, _ := utils.GeneratePassword()

	paramStore := entity.TableMstUser{
		RoleId:    requestModel.RoleId,
		Fullname:  requestModel.Fullname,
		Username:  requestModel.Username,
		IsGender:  requestModel.IsGender,
		Address:   requestModel.Address,
		HpNumber:  requestModel.HpNumber,
		Email:     requestModel.Email,
		Password:  genPass,
		CreatedBy: requestModel.CreatedBy,
		IsActive:  constant.ACTIVE,
	}

	_ = controller.UserRepository.Insert(c.Context(), paramStore)

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       constant.SUCCESSFULLY_ADD,
	})
}

func (controller UserController) Update(c *fiber.Ctx) error {
	var requestModel criteria.UpdateUserCriteria

	err := c.BodyParser(&requestModel)
	exception.PanicLogging(err)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
	}

	idStr := c.Params("id")
	validation.ValidateCriteria(requestModel)

	paramStore := entity.TableMstUser{
		Id:        idStr,
		RoleId:    requestModel.RoleId,
		Fullname:  requestModel.Fullname,
		Username:  requestModel.Username,
		IsGender:  requestModel.IsGender,
		Address:   requestModel.Address,
		HpNumber:  requestModel.HpNumber,
		Email:     requestModel.Email,
		UpdatedBy: requestModel.CreatedBy,
	}
	_, err = controller.UserRepository.FindById(c.Context(), paramStore.Id)
	if err != nil {
		utils.NewLogger().Info(constant.DATA_NOT_FOUND, ": ", err.Error())
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	_ = controller.UserRepository.Update(c.Context(), paramStore)
	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       constant.SUCCESSFULLY_UPDATE,
	})
}

func (controller UserController) UpdateIsActive(c *fiber.Ctx) error {
	var requestModel criteria.IsActiveCriteria

	err := c.BodyParser(&requestModel)
	exception.PanicLogging(err)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
	}

	idStr := c.Params("id")
	validation.ValidateCriteria(requestModel)

	paramStore := entity.TableMstUser{
		Id:        idStr,
		IsActive:  requestModel.IsActive,
		UpdatedBy: requestModel.UpdatedBy,
	}

	_, err = controller.UserRepository.FindById(c.Context(), paramStore.Id)
	if err != nil {
		utils.NewLogger().Info(constant.DATA_NOT_FOUND, ": ", err.Error())
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	_ = controller.UserRepository.Update(c.Context(), paramStore)
	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       constant.SUCCESSFULLY_DELETE,
	})
}

func (controller UserController) GetAllDataUserByParam(c *fiber.Ctx) error {
	var requestModel criteria.UserSearchCriteria

	err := c.BodyParser(&requestModel)
	exception.PanicLogging(err)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
	}

	validation.ValidateCriteria(requestModel)

	resultData, err := controller.UserRepository.FindByParam(c.Context(), requestModel.Key, requestModel.Value)
	if err != nil {
		utils.NewLogger().Info(constant.DATA_NOT_FOUND, ": ", err.Error())
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       resultData,
	})
}

func (controller UserController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	resultData, err := controller.UserRepository.FindById(c.Context(), id)
	if err != nil {
		utils.NewLogger().Info(constant.DATA_NOT_FOUND, ": ", err.Error())
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       resultData,
	})
}

func (controller UserController) FindAll(c *fiber.Ctx) error {

	paging := paginate.PreparePagination(map[string]string{
		"search":         c.Query("search"),
		"sort_by":        c.Query("sort_by"),
		"sort_direction": c.Query("sort_direction"),
		"limit":          c.Query("limit"),
		"page":           c.Query("page"),
	}, []string{
		"id",
		"created_at",
		"updated_at",
	})

	searchByQuery := c.Query("search_by")
	statusByQUery := c.Query("status")
	status := []string{}
	if statusByQUery != "" {
		status = strings.Split(statusByQUery, ",")
	}

	options := criteria.GetListOfOptions{
		SearchBy: searchByQuery,
		Status:   status,
	}

	totalCount, resultData, err := controller.UserRepository.FindAll(c.Context(), paging, options)

	if err != nil {
		utils.NewLogger().Info(constant.DATA_NOT_FOUND, ": ", err.Error())
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	data := result.DataPagingResult{
		PageNumber:       paging.Page,
		PageSize:         paging.Limit,
		TotalRecordCount: totalCount,
		Records:          resultData,
	}

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       data,
	})
}
