package user

// kolom pendaftaran yang fieldnya sesuai dengan FE dan dimapping ke struct User
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"` //using format email
	Password   string `json:"password" binding:"required"`
}
