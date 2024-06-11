package router

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

const JWT_SECRET = "asfbgakjl;gawobi;ioragjewgnVENBVRWOB"

func connectToDB(c *gin.Context) *pgx.Conn {
	connStr := "postgres://postgres:warehouse-manager@localhost:5432/postgres?sslmode=disable"
	conn, err := pgx.Connect(c.Request.Context(), connStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error connecting to the database",
		})
		return nil
	}
	return conn
}

func hashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return bytes, err
}

func checkPasswordValidity(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func createJwtToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"expire": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return ""
	}

	return tokenString
}

func checkTokenAlg(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return JWT_SECRET, nil
}

func parseJwtToken(tokenString string) (string, float64) {
	token, err := jwt.Parse(tokenString, checkTokenAlg)
	if err != nil {
		return "", 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId := claims["userId"].(string)
		expire := claims["expire"].(float64)
		return userId, expire
	} else {
		return "", 0
	}
}

func authMiddleware(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"][0]

	if authHeader == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	token := strings.Split(authHeader, " ")[0]

	userId, expTimestamp := parseJwtToken(token)

	if expTimestamp < float64(time.Now().Unix()) {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.Set("userEmail", userId)
	c.Next()
}
