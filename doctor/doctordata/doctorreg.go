package doctor

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type DoctorReg struct {
	DoctorID *int   `json:"doctorId"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
}

type Login struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
