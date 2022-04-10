package controller

import (
	"database/sql"
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

	db := c.Keys["db"].(*sql.DB)
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Error in reading body"})
		return
	}

	var sqlStatement = "SELECT email FROM users WHERE email=$1"
	rs := db.QueryRow(sqlStatement, user.Email)

	var email string

	_ = rs.Scan(&email)

	if email != "" {
		c.JSON(http.StatusOK, gin.H{"Message": "User email exist"})
		return
	}

	sqlStatement = "SELECT name FROM users WHERE name=$1"
	rs = db.QueryRow(sqlStatement, user.Name)

	var name string

	_ = rs.Scan(&name)

	if name != "" {
		c.JSON(http.StatusOK, gin.H{"Message": "User name exist"})
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	sqlStatement = `INSERT INTO users (
			name,
			email,
			password,
			role
		)	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err = db.QueryRow(sqlStatement,
		user.Name, user.Email, user.Password, "user").Scan(&user.ID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func SignIn(c *gin.Context) {

	var authdetails models.Authentication
	err := c.BindJSON(&authdetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var authuser models.User

	var sqlStatement = `SELECT
		id, name, email, password, role
	FROM users
	WHERE email=$1`

	db := c.Keys["db"].(*sql.DB)
	rs := db.QueryRow(sqlStatement, authdetails.Email)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var token models.Token
	token.Email = authuser.Email
	token.Role = authuser.Role
	token.TokenString = validToken
	c.JSON(http.StatusOK, token)
}

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("Token") == "" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Not allowed"})
			c.Abort()
			return
		}

		secretkey := os.Getenv("SECRET_KEY")
		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(c.Request.Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Token has been expired"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println(claims["email"])
			role := claims["role"]
			c.Set("Role", role)
			c.Next()
			return
		}
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"message": "Authentication Required"})
		c.Abort()
	}
}
