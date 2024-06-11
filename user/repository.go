package user

import "gorm.io/gorm"

// penyimpanan function yang akan dipanggil handler
type Repository interface {
	Save(user User) (User, error)           //register user
	FindByEmail(email string) (User, error) //login
}

// config db
type repository struct {
	db *gorm.DB
}

// membuat koneksi baru untuk penyimpanan data
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// menyimpan data user ke db
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

// mencocokan email & pass
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}
