package user

import (
	"latihan-assesment-radika/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entities.User, error)
	Create(user entities.User) (entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
