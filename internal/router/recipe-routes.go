package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddRecipeRouter(rg *gin.RouterGroup) {
	recipeRouter := rg.Group("/recipe")
	recipeRouter.Use(authMiddleware)
	recipeRouter.GET("/", getRecipes)
	recipeRouter.GET("/:id", getRecipe)
	recipeRouter.POST("/", createRecipe)
	recipeRouter.PUT("/:id", updateRecipe)
	recipeRouter.DELETE("/:id", deleteRecipe)
	recipeRouter.GET("/material/:materialId", getAllRecipesByMaterialId)
}

func getRecipes(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	recipes, err := dbConn.GetAllRecipes(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if recipes == nil {
		recipes = []db.Recipe{}
	}
	c.JSON(200, recipes)
}

func getRecipe(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	recipe, err := dbConn.GetRecipeById(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, recipe)
}

func createRecipe(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	recipeBody := db.CreateRecipeParams{}
	err = json.Unmarshal(bodyStr, &recipeBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateRecipe(c.Request.Context(), recipeBody)
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

	recipeBody := db.UpdateRecipeParams{}
	err = json.Unmarshal(bodyStr, &recipeBody)
	if err != nil {
		c.Status(500)
	}
	recipeBody.ID = []byte(uuid.New().String())

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateRecipe(c.Request.Context(), recipeBody)
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

func getAllRecipesByMaterialId(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	recipes, err := dbConn.GetRecipesByMaterial(c.Request.Context(), []byte(c.Param("materialId")))
	if err != nil {
		c.Status(500)
		return
	}

	if recipes == nil {
		recipes = []db.Recipe{}
	}
	c.JSON(200, recipes)
}
