package pkg

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rishad004/learning-platform-go/user-service/internal/model"
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

	Db.AutoMigrate(&model.Users{})
	return Db
}
