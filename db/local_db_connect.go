package db

import (
	"golang-docker-demo/config"
	"golang-docker-demo/db/generated/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewLocalQuery(cfg *config.Config) *query.Query {
	return query.Use(NewLocalDb(cfg))
}

func NewLocalDb(config *config.Config) *gorm.DB {
	dsn := "host=fiber_postgres user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// db.AutoMigrate(&model.Pet{})
	return db
}
