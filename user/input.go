package user

// KUMPULAN STRUCT YANG AKAN DIPASSING SERVICE

// kolom pendaftaran yang fieldnya sesuai dengan FE dan dimapping ke struct User
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"` //using format email
	Password   string `json:"password" binding:"required"`
}

// input login pada bussiness logic yang sesuai di FE
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
