package routes

import (
	"go-backend/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PublicRoutes(a *fiber.App, dbClient *gorm.DB) {

	//wire controller
	authController := controllers.NewAuthController(dbClient)

	// Empliyee Routes
	v1 := a.Group("/api/v1")

	// Auth Routes
	auth_routes := v1.Group("/auth")
	auth_routes.Post("/login", authController.Login)
	auth_routes.Post("/register", authController.Register)
}
