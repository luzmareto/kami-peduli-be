package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Bussiness logic
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error) // pengecekan avail email saat registrasi
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserById(ID int) (User, error) //bisa digunakan untuk token
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
		return user, errors.New("no user found thath email")
	}

	// pencocokan password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

// pengecekan emaail saat register
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	// pengecekan email yang sudah terdaftar = GAGAL
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// pengecekan email yang belum terdaftar. Berhasil
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// mencari id yang akan melakukan update avatar
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	// menyimpan perubahan ke db
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserById(ID int) (User, error) {
	// memanggil FindByEmail di repository.go untuk melakukan pencocokan
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that id")
	}

	return user, nil
}
