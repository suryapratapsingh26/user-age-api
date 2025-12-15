package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/surya/user-age-api/config"
	"github.com/surya/user-age-api/internal/handler"
	"github.com/surya/user-age-api/internal/logger"
	"github.com/surya/user-age-api/internal/repository"
	"github.com/surya/user-age-api/internal/routes"
	"github.com/surya/user-age-api/internal/service"
)

func main() {
	// Initialize logger
	logger.Init()
	defer logger.Log.Sync()

	app := fiber.New()

	db, err := config.NewDB()
	if err != nil {
		logger.Log.Fatal("failed to connect to database")
	}
	defer db.Close()

	logger.Log.Info("database connection established")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	routes.RegisterUserRoutes(app, userHandler)

	logger.Log.Info("server starting on port 3000")
	log.Fatal(app.Listen(":3000"))
}
