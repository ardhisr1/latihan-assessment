package entities

import "time"

type Book struct {
	ID         int       `gorm:"primaryKey" json:"book_id"`
	UserID     int       `json:"user_id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Year       int       `json:"year"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
