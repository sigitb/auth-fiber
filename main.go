package main

import (
	"base-fiber/app/configs"
	"base-fiber/app/database"
	"base-fiber/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDatabaseConnection()
)

func main() {
	defer database.CloseDatabase(db)


	app := fiber.New()
	app.Static("/storage", configs.ProjectRootPath + "/storage")
	routes.InitApiV1Routers(app, db)

	app.Listen(":" + os.Getenv("GO_PORT"))
}