package main

import (
	"encoding/json"
	"log"
	"net/http"
	"user/database"
	"user/entities"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const SECRET_KEY = `your-256-bit-secret`

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

type LoginResponse struct {
	Response string `json:"response,omitempty"`
	Message  string `json:"message,omitempty"`
	Token    string `json:"token,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	var dbUser entities.User

	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		resp := LoginResponse{Response: err.Error()}
		sendWrappedResponse(w, http.StatusUnprocessableEntity, resp)
		return
	}

	database.Instance.Where("username = ?", user.Username).First(&dbUser)
	if dbUser.Username == "" {
		resp := LoginResponse{Response: "Wrong Username or Password!"}
		sendWrappedResponse(w, http.StatusUnprocessableEntity, resp)
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		resp := LoginResponse{Response: "Wrong Username or Password!"}
		sendWrappedResponse(w, http.StatusForbidden, resp)
		return
	}

	jwtToken, err := GenerateJWT()
	if err != nil {
		resp := LoginResponse{Message: err.Error()}
		sendWrappedResponse(w, http.StatusInternalServerError, resp)
		return
	}

	resp := LoginResponse{Token: jwtToken}
	sendWrappedResponse(w, http.StatusOK, resp)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		sendWrappedResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user.Password = getHash([]byte(user.Password))

	database.Instance.Create(&user)
	user.Password = "*****"
	sendWrappedResponse(w, http.StatusCreated, user)
}

func sendWrappedResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
