package db

import (
	"golang-docker-demo/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewLocalDb(config config.DBConfig) *gorm.DB {
	dsn := "host=fiber_postgres user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// db.AutoMigrate(&model.Pet{})
	return db
}
