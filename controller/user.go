package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"fyc.com/sprs/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
		_ = fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func SignUp(c *gin.Context) {

	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			ID:      1,
			Message: "Error in reading body",
		})
		return
	}

	var sqlStatement = "SELECT email FROM users WHERE email=$1"
	rs := Sql().QueryRow(sqlStatement, user.Email)

	var email string

	_ = rs.Scan(&email)

	if email != "" {
		c.JSON(http.StatusOK, Response{
			ID:      1,
			Message: "User email exist",
		})
		return
	}

	sqlStatement = "SELECT name FROM users WHERE name=$1"
	rs = Sql().QueryRow(sqlStatement, user.Email)

	var name string

	_ = rs.Scan(&name)

	if name != "" {
		c.JSON(http.StatusOK, Response{
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
		c.JSON(http.StatusOK, Response{
			ID:      1,
			Message: "Error creating user",
		})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func SignIn(c *gin.Context) {

	var authdetails models.Authentication
	err := c.BindJSON(&authdetails)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	check := CheckPasswordHash(authdetails.Password, authuser.Password)

	if !check {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	validToken, err := GenerateJWT(authuser.Email, authuser.Role)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var token models.Token
	token.Email = authuser.Email
	token.Role = authuser.Role
	token.TokenString = validToken
	//w.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, token)
}

// func IsAuthorized(c *gin.Context) http.HandlerFunc {
// 	return func(c *gin.Context) {

// 		if r.Header["Token"] == nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		secretkey := os.Getenv("SECRET_KEY")
// 		var mySigningKey = []byte(secretkey)

// 		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("there was an error in parsing")
// 			}
// 			return mySigningKey, nil
// 		})

// 		if err != nil {
// 			http.Error(w, "Token has been expired", 404)
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			if claims["role"] == "admin" {

// 				r.Header.Set("Role", "admin")
// 				handler.ServeHTTP(w, r)
// 				return

// 			} else if claims["role"] == "user" {

// 				r.Header.Set("Role", "user")
// 				handler.ServeHTTP(w, r)
// 				return
// 			}
// 		}
// 		http.Error(w, "Not Authorized", 404)
// 		c.JSON(http.StatusOK, err)
// 	}
