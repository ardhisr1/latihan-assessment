package entities

import "time"

type User struct {
	ID         int       `gorm:"primaryKey" json:"user_id"`
	Name       string    `json:"name"`
	DateBirth  time.Time `json:"date_birth"`
	Address    string    `json:"address"`
	Email      string    `gorm:"index:idx_email, unique" json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	//Books      []Book    `gorm:"foreignKey:UserID"`
}

type UserInput struct {
	Name      string    `json:"name" binding:"required"`
	DateBirth time.Time `json:"date_birth" binding:"required"`
	Address   string    `json:"address"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
}

type UserUpdateInput struct {
	Name      string    `json:"name" binding:"required"`
	DateBirth time.Time `json:"date_birth" binding:"required"`
	Address   string    `json:"address"`
}
