package main

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

type PasswordGenerated struct {
	gorm.Model
	Key      string `json:"key,omitempty"`
	UserName string `json:"username,omitempty"`
	Value    string `json:"value,omitempty"`
}

var DB Dbinstance

// connectDb
func ConnectDb() {

	db, err := gorm.Open(sqlite.Open("passwords.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	db.Logger = logger.Default.LogMode(logger.Silent)

	err = db.AutoMigrate(&PasswordGenerated{})
	if err != nil {
		return
	}
	//	DB.AutoMigrate(&model.Product{}, &model.User{})
	DB = Dbinstance{
		Db: db,
	}
}
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
