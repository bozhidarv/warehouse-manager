package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddUnitRouter(rg *gin.RouterGroup) {
	supplierRouter := rg.Group("/units")
	supplierRouter.GET("/", getUnits)
	supplierRouter.GET("/:id", getUnit)
	supplierRouter.POST("/", createUnit)
	supplierRouter.PUT("/:id", updateUnit)
	supplierRouter.DELETE("/:id", deleteUnit)
}

func getUnits(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	units, err := dbConn.GetAllUnits(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if units == nil {
		units = []db.Unit{}
	}
	c.JSON(200, units)
}

func getUnit(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	unit, err := dbConn.GetUnitById(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, unit)
}

func createUnit(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	unitBody := db.CreateUnitParams{}
	err = json.Unmarshal(bodyStr, &unitBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateUnit(c.Request.Context(), unitBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func updateUnit(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	unitBody := db.UpdateUnitParams{}
	err = json.Unmarshal(bodyStr, &unitBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateUnit(c.Request.Context(), unitBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteUnit(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteUnit(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
