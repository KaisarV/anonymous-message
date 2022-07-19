package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	model "ngl-link/models"
	"ngl-link/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Jksdgbfkd334dsj")
var tokenName = "token"

func generateToken(w http.ResponseWriter, id int, name string, username string, email string, userType int) {
	timeout, _ := strconv.Atoi(utils.Getenv("TOKEN_MINUTE_LIFESPAN", "1"))

	tokenExpiryTime := time.Now().Add(time.Duration(timeout) * time.Minute)

	claims := &model.Claims{
		ID:       id,
		Name:     name,
		Username: username,
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}
func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken, userType := validateUserToken(w, r, accessType)
		if !isValidToken {
			if userType == -1 {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
			SendUnAuthorizedResponse(w, "Unauthorized Response", 400)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(w http.ResponseWriter, r *http.Request, accessType int) (bool, int) {
	isAccessTokenValid, id, name, username, email, userType := validateTokenFromCookies(r)
	fmt.Print(id, name, username, email, userType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType >= accessType
		fmt.Print(isUserValid)
		if isUserValid {
			return true, userType
		}
	}
	return false, userType
}

//Ambil dari cookies
func validateTokenFromCookies(r *http.Request) (bool, int, string, string, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &model.Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.Username, accessClaims.Email, accessClaims.UserType
		}
	}
	return false, -1, "", "", "", -1
}

func GetDataFromToken(token string) int {

	claims := &model.Claims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(accessToken *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == nil && parsedToken.Valid {
		return claims.ID
	} else {
		return -1
	}

}

func SendUnAuthorizedResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	var response model.MessageResponse
	response.Status = status
	response.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
