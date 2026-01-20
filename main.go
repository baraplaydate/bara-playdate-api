package main

import (
	_ "bara-playdate-api/docs"
	"bara-playdate-api/exception"
	"bara-playdate-api/utils"

	controller "bara-playdate-api/api/controllers"
	implRepository "bara-playdate-api/repository/impl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {

	//setup configuration
	config := utils.NewEnv()
	database := utils.NewDatabase(config)
	// redis := configuration.NewRedis(config)

	//repository
	authRepository := implRepository.NewAuthRepositoryImpl(database)
	roleRepository := implRepository.NewRoleRepositoryImpl(database)
	userRepository := implRepository.NewUserRepositoryImpl(database)

	//controller
	authController := controller.NewAuthController(&authRepository, config)
	roleController := controller.NewRoleController(&roleRepository, config)
	userController := controller.NewUserController(&userRepository, config)
	commonController := controller.NewCommonController(config)
	httpBinController := controller.NewHttpBinController()

	//setup fiber
	app := fiber.New(utils.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	authController.Route(app)
	roleController.Route(app)
	userController.Route(app)
	commonController.Route(app)
	httpBinController.Route(app)

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//start app
	err := app.Listen(config.ServerPort)
	exception.PanicLogging(err)
}
