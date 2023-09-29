package controllers

import (
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

type UserDTO struct {
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
}

func (uc *UserController) GetOwnProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	
	var ownProfil models.User
	if err := uc.DB.Find(&ownProfil, "email = ?", email).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": err.Error(),
		})
	}

	userDTO := UserDTO{
		Name: ownProfil.Name,
		Email: ownProfil.Email,
		PhoneNumber: ownProfil.PhoneNumber,
		Address: ownProfil.Address,
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Profil Anda berhasil diambil",
		"data": fiber.Map{"user": userDTO},
	})
} 

