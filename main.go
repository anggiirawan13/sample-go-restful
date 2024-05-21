package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sample-go-restful/controllers"
	"github.com/sample-go-restful/models"
	"github.com/sample-go-restful/repositories"
	"github.com/sample-go-restful/routers"
	"github.com/sample-go-restful/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=130302 dbname=restful port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		return
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	app := fiber.New()

	routers.SetupRoutes(app, userController)

	err = app.Listen("localhost:3000")
	if err != nil {
		return
	}
}
