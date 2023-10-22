package database

import (
	"github.com/devmarciosieto/Client-Server-API/server/internal/domain/entity"
	"gorm.io/gorm"
)

type USDBRLRepository struct {
	Db *gorm.DB
}

func NewUSDBRLRepository(db *gorm.DB) *USDBRLRepository {
	return &USDBRLRepository{Db: db}
}

func (u *USDBRLRepository) InsertUSDBRL(usdbrl *entity.USDBRL) error {
	result := u.Db.Create(usdbrl)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
