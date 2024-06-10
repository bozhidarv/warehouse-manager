package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddRecipeRouter(rg *gin.RouterGroup) {
	recipeRouter := rg.Group("/recipe")
	recipeRouter.GET("/", getRecipes)
	recipeRouter.GET("/:id", getRecipe)
	recipeRouter.POST("/", createRecipe)
	recipeRouter.PUT("/:id", updateRecipe)
	recipeRouter.DELETE("/:id", deleteRecipe)
}

func getRecipes(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	units, err := dbConn.GetAllRecipes(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if units == nil {
		units = []db.Recipe{}
	}
	c.JSON(200, units)
}

func getRecipe(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	unit, err := dbConn.GetRecipeById(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, unit)
}

func createRecipe(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	unitBody := db.CreateRecipeParams{}
	err = json.Unmarshal(bodyStr, &unitBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateRecipe(c.Request.Context(), unitBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func updateRecipe(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	unitBody := db.UpdateRecipeParams{}
	err = json.Unmarshal(bodyStr, &unitBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateRecipe(c.Request.Context(), unitBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteRecipe(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteRecipe(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
