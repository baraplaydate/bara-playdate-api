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

type RoleController struct {
	repository.RoleRepository
	utils.Config
}

func NewRoleController(roleRepository *repository.RoleRepository, config utils.Config) *RoleController {
	return &RoleController{RoleRepository: *roleRepository, Config: config}
}

func (controller RoleController) Route(app *fiber.App) {

	group := app.Group(controller.Config.Route)

	group.Post("/role/store", middleware.AuthenticationJWT(controller.Config), controller.Create)
	group.Put("/role/update/:id", middleware.AuthenticationJWT(controller.Config), controller.Update)
	group.Put("/role/updateIsActive/:id", middleware.AuthenticationJWT(controller.Config), controller.UpdateIsActive)
	group.Get("/role/getDataById/:id", middleware.AuthenticationJWT(controller.Config), controller.FindById)
	group.Get("/role", middleware.AuthenticationJWT(controller.Config), controller.FindAll)
}

func (controller RoleController) Create(c *fiber.Ctx) error {
	var requestModel criteria.StoreRoleCriteria

	err := c.BodyParser(&requestModel)
	exception.PanicLogging(err)

	validation.ValidateCriteria(requestModel)

	paramStore := entity.TableMstAccRole{
		RoleName:    requestModel.RoleName,
		Description: requestModel.Description,
		CreatedBy:   requestModel.CreatedBy,
		IsActive:    constant.ACTIVE,
	}

	_ = controller.RoleRepository.Insert(c.Context(), paramStore)

	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       constant.SUCCESSFULLY_ADD,
	})
}

func (controller RoleController) Update(c *fiber.Ctx) error {
	var requestModel criteria.UpdateRoleCriteria

	err := c.BodyParser(&requestModel)
	exception.PanicLogging(err)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
	}

	idStr := c.Params("id")
	validation.ValidateCriteria(requestModel)

	paramStore := entity.TableMstAccRole{
		Id:          idStr,
		RoleName:    requestModel.RoleName,
		Description: requestModel.Description,
		UpdatedBy:   requestModel.UpdatedBy,
	}

	_, err = controller.RoleRepository.FindById(c.Context(), paramStore.Id)
	if err != nil {
		utils.NewLogger().Info(constant.DATA_NOT_FOUND, err.Error())
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	_ = controller.RoleRepository.Update(c.Context(), paramStore)
	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       constant.SUCCESSFULLY_UPDATE,
	})
}

func (controller RoleController) UpdateIsActive(c *fiber.Ctx) error {
	var requestModel criteria.IsActiveCriteria

	err := c.BodyParser(&requestModel)
	exception.PanicLogging(err)
	if err != nil {
		utils.NewLogger().Info(constant.ERR_OBJECT_VALIDATION_DETAIL, ": ", err.Error())
	}

	idStr := c.Params("id")
	validation.ValidateCriteria(requestModel)

	paramStore := entity.TableMstAccRole{
		Id:        idStr,
		IsActive:  requestModel.IsActive,
		UpdatedBy: requestModel.UpdatedBy,
	}

	_, err = controller.RoleRepository.FindById(c.Context(), paramStore.Id)
	if err != nil {
		utils.NewLogger().Info(constant.DATA_NOT_FOUND, ": ", err.Error())
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	_ = controller.RoleRepository.Update(c.Context(), paramStore)
	return c.JSON(result.ResponseResult{
		ResponseCode:        http.StatusOK,
		ResponseDescription: constant.SUCCESS,
		ResponseTime:        utils.DateToStdNow(),
		ResponseDatas:       constant.SUCCESSFULLY_DELETE,
	})
}

func (controller RoleController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	resultData, err := controller.RoleRepository.FindById(c.Context(), id)
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

func (controller RoleController) FindAll(c *fiber.Ctx) error {

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

	totalCount, resultData, err := controller.RoleRepository.FindAll(c.Context(), paging, options)

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
