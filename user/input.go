package user

// struct untuk input dari form
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"Occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
