package user

import "gorm.io/gorm"

// penyimpanan function yang akan dipanggil handler
type Repository interface {
	Save(user User) (User, error)           //register user
	FindByEmail(email string) (User, error) //login
	FindByID(ID int) (User, error)          //bisa untuk apload avatar
	Update(user User) (User, error)         //bisa untuk apload avatar
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

// mencocokan email UNTUK LOGIN
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

// MENCARI ID USER UNTUK UPLOAD AVATAR
func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

// update avatar
func (r *repository) Update(user User) (User, error) {
	// data yang sudah ada akan disimpan
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}
