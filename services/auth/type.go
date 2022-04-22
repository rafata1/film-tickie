package auth

type CheckOTPRequest struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

type CheckOTPResponseData struct {
	Token string `json:"token"`
}

type BaseRes struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UpdateUserInfoRequest struct {
	Name string `json:"name"`
}
