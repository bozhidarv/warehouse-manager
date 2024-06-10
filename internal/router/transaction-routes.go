package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddTransactionHistoryRouter(rg *gin.RouterGroup) {
	supplierRouter := rg.Group("/units")
	supplierRouter.GET("/", getTransactionHistorys)
	supplierRouter.GET("/:id", getTransactionHistory)
	supplierRouter.POST("/", createTransactionHistory)
	supplierRouter.PUT("/:id", updateTransactionHistory)
	supplierRouter.DELETE("/:id", deleteTransactionHistory)
}

func getTransactionHistorys(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	units, err := dbConn.GetAllTransactionHistory(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if units == nil {
		units = []db.TransactionHistory{}
	}
	c.JSON(200, units)
}

func getTransactionHistory(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	unit, err := dbConn.GetTransactionHistoryById(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, unit)
}

func createTransactionHistory(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	unitBody := db.CreateTransactionHistoryParams{}
	err = json.Unmarshal(bodyStr, &unitBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateTransactionHistory(c.Request.Context(), unitBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func updateTransactionHistory(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	unitBody := db.UpdateTransactionHistoryParams{}
	err = json.Unmarshal(bodyStr, &unitBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateTransactionHistory(c.Request.Context(), unitBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteTransactionHistory(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteTransactionHistory(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
