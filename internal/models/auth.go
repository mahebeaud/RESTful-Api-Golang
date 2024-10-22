package models

// POST Forms for login
type AuthLogin struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

// POST Forms for register
type AuthRegister struct {
	Email           string `binding:"required" json:"email"`
	Username        string `binding:"required" json:"username"`
	Password        string `binding:"required" json:"password"`
	ConfirmPassword string `binding:"required" json:"confirm_password"`
}

// POST Forms for register
type AuthLogout struct {
	Uuid string `binding:"required" json:"uuid"`
}

// Token data Struct
type TokenDataStruct struct {
	Token string
	Uuid string
	Exp string
}