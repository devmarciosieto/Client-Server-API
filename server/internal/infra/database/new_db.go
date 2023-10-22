package database

import (
	"github.com/devmarciosieto/Client-Server-API/server/internal/domain/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("catacao.db"), &gorm.Config{})

	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&entity.USDBRL{})

	return db
}
