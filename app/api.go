package app

import (
	"log"
	"go-backend/app/routes"
	"go-backend/config"
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func StartAppApi() {
	if err := config.InitialDB(); err != nil {
		panic("koneksi database gagal")
	}
	
	//daftar model yang akan dimigrate
	modeslToMigrate := []interface{}{
		&models.User{},
	}

	// Cek apakah migrasi sudah dijalankan sebelumnya
	for _, model := range modeslToMigrate {
		if !config.DB.Migrator().HasTable(model) {
			// Jika belum ada tabel, jalankan AutoMigrate
			config.DB.AutoMigrate(model)
		}
	}

	// Start App
	app := fiber.New()

	// Cors and Logger
	app.Use(cors.New())
	app.Use(fiberLogger.New())

	// Init Routes
	routes.PublicRoutes(app, config.DB)
	routes.PrivateRoutes(app, config.DB)

	log.Fatal(app.Listen("localhost:3000"))

}

