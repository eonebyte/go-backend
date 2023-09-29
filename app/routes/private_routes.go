package routes

import (
	"os"
	"go-backend/controllers"
	"go-backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PrivateRoutes(a *fiber.App, dbClient *gorm.DB) {

	//Middlewares
	secret_key := os.Getenv("SECRET_KEY")
	jwtAuth := middlewares.NewAuthMiddleware(secret_key)

	//wire controller
	userController := controllers.NewUserController(dbClient)

	// Init Routes
	v1 := a.Group("/api/v1")

	// User Routes
	user_routes := v1.Group("/users")
	user_routes.Get("/me", jwtAuth, userController.GetOwnProfile)
}
