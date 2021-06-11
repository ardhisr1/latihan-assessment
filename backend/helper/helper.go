package helper

import (
	"errors"
	"latihan-assesment-radika/entities"
	"math"
	"strconv"
	"time"
)

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

func ValidateID(ID string) error {
	if num, err := strconv.Atoi(ID); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("Input must be number")
	}
	return nil
}
