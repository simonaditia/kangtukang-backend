package models

type AuthenticationInputRegister struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthenticationInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
