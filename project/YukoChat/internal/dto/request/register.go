package request

type RegisterRequest struct {
	Telephone        string `json:"telephone" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Nickname         string `json:"nickname" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verification_code" binding:"required"`
}
