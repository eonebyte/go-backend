package controllers

import (
	"os"
	"time"
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

type NewUserDTO struct {
	Name string `json:"name"`
	Email string `jsonL:"email"`
}

type RequestLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		DB: db,
	}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {

	var reqLogin RequestLogin
	
	if err := c.BodyParser(&reqLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "request gagal",
			"message": "Gagal memuat request login",
		})
	}

	var users []models.User
	if err := ac.DB.Find(&users).Error; err != nil {
		errorMessage := "Data pengguna tidak ditemukan" + err.Error()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "not found",
			"code":    404,
			"message": errorMessage,
		})
	}

	// Throws Unauthorized error
	for _, user := range users {
		if reqLogin.Email == user.Email && comparePassword([]byte(user.Password), reqLogin.Password) {

			// Create the Claims
			claims := jwt.MapClaims{
				"name":  "Iwan Byte",
				"email": user.Email,
				"admin": true,
				"exp":   time.Now().Add(time.Hour * 72).Unix(),
			}

			// Create token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			secret_key := os.Getenv("SECRET_KEY")
			t, err := token.SignedString([]byte(secret_key))
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			return c.JSON(fiber.Map{"status": "success","token": t})
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "unauthorized",
		"message": "Email atau kata sandi salah",
	})
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	newUser := new(models.User)

	if err := c.BodyParser(newUser); err != nil {
		errorMessage := "Gagal memuat permintaan" + err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "gagal",
			"message": errorMessage,
		})
	}
	hashedPass := hashPassword(newUser.Password)
	newUser.Password = string(hashedPass)

	if result := ac.DB.Create(newUser); result.Error != nil {
		errorMessage :=  "Gagal membuat data user" + result.Error.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "gagal",
			"message":  errorMessage,
		})
	}

	newUserDTO := NewUserDTO{
		Name: newUser.Name,
		Email: newUser.Email,
	}


	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Registrasi berhasil",
		"data":    fiber.Map{"newUser": newUserDTO},
	})
}

func hashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func comparePassword(hashedPassword []byte, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword))
	return err == nil
}
