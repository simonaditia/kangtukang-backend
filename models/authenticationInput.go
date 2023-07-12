package models

type AuthenticationInputRegister struct {
	Nama      string  `json:"nama" binding:"required"`
	Email     string  `json:"email"`
	Password  string  `json:"password" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	NoTelp    string  `json:"no_telp"`
}

type AuthenticationInput struct {
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
	NoTelp   string `json:"no_telp"`
}
