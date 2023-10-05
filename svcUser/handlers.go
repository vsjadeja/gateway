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

func Login(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	var dbUser entities.User

	json.NewDecoder(r.Body).Decode(&user)

	if err := user.Validate(); err != nil {
		sendWrappedResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	database.Instance.Where("username = ?", user.Username).First(&dbUser)
	if dbUser.Username == "" {
		sendWrappedResponse(w, http.StatusUnprocessableEntity, `{"response":"Wrong Username or Password!"}`)
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		sendWrappedResponse(w, http.StatusOK, `{"response":"Wrong Username or Password!"}`)
		return
	}

	jwtToken, err := GenerateJWT()
	if err != nil {
		sendWrappedResponse(w, http.StatusInternalServerError, `{"message":"`+err.Error()+`"}`)
		return
	}

	sendWrappedResponse(w, http.StatusOK, `{"token":"`+jwtToken+`"}`)
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
