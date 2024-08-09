package pkg

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	Db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")))
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate()
	return Db
}
