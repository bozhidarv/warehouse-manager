package router

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddUserRouter(rg *gin.RouterGroup) {
	userRouter := rg.Group("/users")
	userRouter.Use(authMiddleware)
	userRouter.GET("/", getUsers)
	userRouter.GET("/:email", getUser)
	userRouter.PUT("/:email", updateUser)
	userRouter.DELETE("/:email", deleteUser)

	authRouter := rg.Group("")
	authRouter.POST("/login", login)
	authRouter.POST("/register", register)
	authRouter.POST("/logout", logout)
}

func getUsers(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	users, err := dbConn.GetAllUsers(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if users == nil {
		users = []db.User{}
	}
	c.JSON(200, users)
}

func getUser(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	user, err := dbConn.GetUserByEmail(c.Request.Context(), c.Param("email"))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, user)
}

func updateUser(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	userBody := db.UpdateUserParams{}
	err = json.Unmarshal(bodyStr, &userBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateUser(c.Request.Context(), userBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteUser(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteUser(c.Request.Context(), c.Param("email"))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}

func login(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	userBody := db.CreateUserParams{}
	err = json.Unmarshal(bodyStr, &userBody)
	if err != nil {
		c.Status(500)
		return
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	if err != nil {
		c.Status(500)
		return
	}

	dbUser, err := dbConn.GetUserByEmail(c, userBody.Email)
	if err != nil {
		c.Status(500)
		return
	}

	if !checkPasswordValidity(userBody.PasswordHash, dbUser.PasswordHash) {
		c.Status(400)
		return
	}

	token := createJwtToken(dbUser.Email)

	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Status(200)
}

func register(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	userBody := db.CreateUserParams{}
	err = json.Unmarshal(bodyStr, &userBody)
	if err != nil {
		c.Status(500)
	}

	userBody.PasswordHash, err = hashPassword(userBody.PasswordHash)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateUser(c.Request.Context(), userBody)
	if err != nil {
		c.Status(500)
	}

	token := createJwtToken(userBody.Email)

	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Status(201)
}

func logout(c *gin.Context) {
	c.Header("Authorization", "")
	c.Status(200)
}
