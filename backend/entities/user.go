package entities

import "time"

type User struct {
	ID         int       `gorm:"primaryKey" json:"user_id"`
	Name       string    `json:"address"`
	DateBirth  time.Time `json:"date_birth"`
	Email      string    `gorm:"index:idx_email, unique" json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Books      []Book    `gorm:"foreignKey:UserID"`
}
