package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string
	Email    string
	UserType int `json:"userType"`
	jwt.StandardClaims
}
