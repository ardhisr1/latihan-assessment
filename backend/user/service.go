package user

import (
	"errors"
	"fmt"
	"latihan-assesment-radika/entities"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUser() (UserFormat, error)
	SaveNewUser(user entities.UserInput) (UserFormat, error)
	UpdateUserByID(id string, input entities.UserUpdateInput) (UserFormat, error)
	DeleteUserBYID(id string) (interface{}, error)
	GetUserByID(id string) (UserFormat, error)
	LoginUser(user entities.LoginUser) (LoginResponseFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	users, err := s.repository.FindAll()

	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormattingUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SaveNewUser(user entities.UserInput) (UserFormat, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	var newUser = entities.user{
		Name:      user.Name,
		DateBirth: user.DateBirth,
		Address:   user.Address,
		Email:     user.Email,
		Password:  encrypt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := s.repository.SaveNewUser(newUser)
	formatUser := FormattingUser(created)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) UpdateUserByID(id string, input entities.UserUpdateInput) (UserFormat, error) {
	var dataUp = map[string]interface{}{}

	user, err := s.repository.FindByID(id)

	if err != nil {
		return UserFormat{}, err
	}

	if input.ID == 0 {
		newError := fmt.Sprintf("user id not found")
		return nil, errors.New(newError)
	}

	if input.Name != "" || len(input.Name) {
		dataUp["name"] = input.Name
	}

	if input.DateBirth != "" || len(input.DateBirth) {
		dataUp["date_birth"] = input.DateBirth
	}

	if input.Address != "" || len(input.Address) {
		dataUp["address"] = input.Address
	}

	dataUp["updated_at"] = time.Now()

	updated, err := s.repository.UpdateById(id, dataUp)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormattingUser(updated)

	return formatUser, nil

}

func (s *service) DeleteUserBYID(id string) (string, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id not found")
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteById(id)

	if err != nil {
		return nil, errors.New("error delete in internal server")
	}
	msg := fmt.Sprintf("user id %s success deleted", id)

	return msg
}

func (s *service) GetUserByID(id string) (UserFormat, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id not found")
		return nil, errors.New(newError)
	}

	formatUser := FormattingUser(user)
	return formatUser, nil
}

func (s *service) LoginUser(user entities.LoginUser) (LoginResponseFormat, error) {
	userGet, err := s.repository.FIndByEmail(user.Email)
	if err != nil {
		return LoginResponseFormat{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userGet.Password), []byte(user.Password)); err != nil {
		return LoginResponseFormat{}, errors.New("Password Invalid")
	}
	formatLogin := FormatingLoginResponse(userGet, "TEST")
	return formatLogin
}
