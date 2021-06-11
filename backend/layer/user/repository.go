package user

import (
	"latihan-assesment-radika/entities"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entities.User, error)
	Create(user entities.User) (entities.User, error)
	UpdateById(id string, dataUp map[string]interface{}) (entities.User, error)
	DeleteById(id string) (string, error)
	FindByID(id string) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entities.User, error) {
	var users []entities.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Create(user entities.User) (entities.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateById(id string, dataUp map[string]interface{}) (entities.User, error) {
	var user entities.User

	if err := r.db.Model(&user).Where("user_id = ?", id).Updates(dataUp).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteById(id string) (string, error) {
	if err := r.db.Where("user_id = ?", id).Delete(&entities.User{}).Error; err != nil {
		return "error", err
	}

	return "succes", nil
}

func (r *repository) FindByID(id string) (entities.User, error) {
	var user entities.User

	if err := r.db.Where("user_id = ", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entities.User, error) {
	var user entities.User

	if err := r.db.Where("email = ", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
