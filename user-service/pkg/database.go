package pkg

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	Db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")))
	if err != nil {
		panic(err)
	}
	return Db
}
