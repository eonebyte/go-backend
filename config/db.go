package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialDB() error {
	if DB != nil {
		return nil
	}
	// Memuat variabel lingkungan dari file .env
    err := godotenv.Load()
    if err != nil {
        panic("Gagal memuat file .env")
    }

    db_username := os.Getenv("DB_USERNAME")
    db_password := os.Getenv("DB_PASSWORD")
    db_name := os.Getenv("DB_NAME")

    dsn := "sqlserver://" + db_username + ":" + db_password + "@localhost:1433?database=" + db_name
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connect db sql server")
	}

	DB = db
	return nil
}