package request

type SendVerificationCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}
