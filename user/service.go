package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Bussiness logic
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

// memasukan Repository ke dalam service
type service struct {
	repository Repository
}

// membuat new service
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// mengambil field User struct sesuai kebutuhan RegisterUserInput
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	// mengubah password menjadi hash
	user.PasswordHash = string(passwordHash)
	// saat user mendaftar, defaultnya adalah user bukan admin
	user.Role = "user"

	// memanggil repository untuk menyimpan data register user
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

// pencocokan email & password
func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	// pengecekan error & succses untuk email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found thath email")
	}

	// pencocokan password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
