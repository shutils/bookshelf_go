package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	dsn := "host=db user=postgres password=postgres dbname=bookshelf port=5432 sslmode=disable"
	var err error
	if Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}
	fmt.Println("db connected.")
	Db.AutoMigrate(&BookEntity{})
	fmt.Println("db migrate.")
}
