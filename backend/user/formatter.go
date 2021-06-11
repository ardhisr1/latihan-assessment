package user

import (
	"latihan-assesment-radika/entities"
	"time"
)

type UserFormat struct {
	ID        int       `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	DateBirth time.Time `json:"date_birth"`
	Address   string    `json:"address"`
}

func FormattingUser(user entities.User) UserFormat {
	UserFormat := UserFormat{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		DateBirth: user.DateBirth,
		Address:   user.Address,
	}

	return UserFormat
}

type LoginResponseFormat struct {
	ID            int       `json:"user_id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	DateBirth     time.Time `json:"date_birth"`
	Address       string    `json:"address"`
	Authorization string
}

func FormatingLoginResponse(user entities.User, token string) LoginResponseFormat {
	LoginResponse := LoginResponseFormat{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		DateBirth:     user.DateBirth,
		Address:       user.Address,
		Authorization: token,
	}

	return LoginResponse
}

type DeleteFormat struct {
	Message string `json:"data`
}

func FormattingDeleteUser(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message: msg,
	}
	return deleteFormat
}
