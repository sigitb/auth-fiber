package routes

import (
	"base-fiber/app/handlers"
	"base-fiber/app/middleware"
	"base-fiber/src/role"
	"base-fiber/src/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitApiV1Routers(r *fiber.App, db *gorm.DB) {
	// role
	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)

	// user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := handlers.NewAuthHandler(userService, roleService)

	v1 := r.Group("api/v1")

	auth := v1.Group("/auth").Use(middleware.Authenticate)
	auth.Post("/register", userController.RegisterUser)
	auth.Post("/login", userController.Login)
}