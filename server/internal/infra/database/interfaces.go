package database

import "github.com/devmarciosieto/Client-Server-API/server/internal/domain/entity"

type USDBRLInterface interface {
	InsertUSDBRL(usdbrl *entity.USDBRL) error
}
