package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"fyc.com/sprs/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"net/http"
)

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email, role string) (string, error) {
	secretkey := os.Getenv("SECRET_KEY")
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		_ = fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			ID:      1,
			Message: "Error in reading body",
		})
		return
	}

	var sqlStatement = "SELECT email FROM users WHERE email=$1"
	rs := Sql().QueryRow(sqlStatement, user.Email)

	var email string

	err = rs.Scan(&email)

	if email != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			ID:      1,
			Message: "User email exist",
		})
		return
	}

	sqlStatement = "SELECT name FROM users WHERE name=$1"
	rs = Sql().QueryRow(sqlStatement, user.Email)

	var name string

	err = rs.Scan(&name)

	if name != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			ID:      1,
			Message: "User name exist",
		})
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	sqlStatement = `INSERT INTO users 
	(name, email, password, role)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err = Sql().QueryRow(sqlStatement,
		user.Name, user.Email, user.Password, "user").Scan(&user.ID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			ID:      1,
			Message: "Error creating user",
		})
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	var authdetails models.Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)

	if err != nil {
		http.Error(w, "Error in reading body", 404)
		return
	}

	var authuser models.User

	var sqlStatement = `SELECT
		id, name, email, password, role
	FROM users
	WHERE email=$1`

	rs := Sql().QueryRow(sqlStatement, authdetails.Email)

	err = rs.Scan(
		&authuser.ID,
		&authuser.Name,
		&authuser.Email,
		&authuser.Password,
		&authuser.Role,
	)

	if err != nil {
		http.Error(w, "Email or password was wrong", 404)
		return
	}

	check := CheckPasswordHash(authdetails.Password, authuser.Password)

	if !check {
		http.Error(w, "Email or password was wrong", 404)
		return
	}

	validToken, err := GenerateJWT(authuser.Email, authuser.Role)

	if err != nil {
		http.Error(w, "Failed to generate token", 404)
		return
	}

	var token models.Token
	token.Email = authuser.Email
	token.Role = authuser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			http.Error(w, "Token not found", 404)
			return
		}

		secretkey := os.Getenv("SECRET_KEY")
		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			http.Error(w, "Token has been expired", 404)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Not Authorized", 404)
		json.NewEncoder(w).Encode(err)
	}
}
