package patient

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type PatientReg struct {
	PatientID   *int   `json:"patientId"`
	Email       string `json:"email"`
	Pass        string `json:"pass"`
	Patientname string `json:"patientname"`
	Mobile      string `json:"mobile"`
	Address     string `json:"address"`
}

type Login struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type Claims struct {
	Patientname string `json:"patientname"`
	jwt.StandardClaims
}
