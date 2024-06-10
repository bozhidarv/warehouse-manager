package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddMaterialRouter(rg *gin.RouterGroup) {
	materrialRouter := rg.Group("/materials")
	materrialRouter.GET("/", getMaterials)
	materrialRouter.GET("/:id", getMaterial)
	materrialRouter.POST("/", createMaterial)
	materrialRouter.PUT("/:id", updateMaterial)
	materrialRouter.DELETE("/:id", deleteMaterial)
}

func getMaterials(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	materials, err := dbConn.GetAllMaterials(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if materials == nil {
		materials = []db.Material{}
	}
	c.JSON(200, materials)
}

func getMaterial(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	material, err := dbConn.GetMaterialById(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, material)
}

func createMaterial(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	materialBody := db.CreateMaterialParams{}
	err = json.Unmarshal(bodyStr, &materialBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateMaterial(c.Request.Context(), materialBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func updateMaterial(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	materialBody := db.UpdateMaterialParams{}
	err = json.Unmarshal(bodyStr, &materialBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateMaterial(c.Request.Context(), materialBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteMaterial(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteMaterial(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
