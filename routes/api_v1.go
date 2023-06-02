package routes

import (
	"base-fiber/app/handlers"
	"base-fiber/src/role"
	"base-fiber/src/user"
	"base-fiber/src/verification"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitApiV1Routers(r *fiber.App, db *gorm.DB) {
	// verification 
	verificationRepository := verification.NewRepository(db)
	verificationService := verification.NewService(verificationRepository)

	// role
	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)

	// user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := handlers.NewAuthHandler(userService, roleService,verificationService)

	v1 := r.Group("api/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", userController.RegisterUser)
	auth.Post("/login", userController.Login)
	auth.Post("/send-otp", userController.SendOtp)
	auth.Post("/verification", userController.Verification)
	auth.Post("/update-password", userController.UpdatePassword)
}