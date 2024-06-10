package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddSupplierRouter(rg *gin.RouterGroup) {
	supplierRouter := rg.Group("/suppliers")
	supplierRouter.GET("/", getSuppliers)
	supplierRouter.GET("/:id", getSupplier)
	supplierRouter.POST("/", createSupplier)
	supplierRouter.PUT("/:id", updateSupplier)
	supplierRouter.DELETE("/:id", deleteSupplier)
}

func getSuppliers(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	suppliers, err := dbConn.GetAllSuppliers(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if suppliers == nil {
		suppliers = []db.Supplier{}
	}
	c.JSON(200, suppliers)
}

func getSupplier(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	supplier, err := dbConn.GetSupplierById(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, supplier)
}

func createSupplier(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	supplierBody := db.CreateSupplierParams{}
	err = json.Unmarshal(bodyStr, &supplierBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateSupplier(c.Request.Context(), supplierBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func updateSupplier(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	supplierBody := db.UpdateSupplierParams{}
	err = json.Unmarshal(bodyStr, &supplierBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateSupplier(c.Request.Context(), supplierBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteSupplier(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteSupplier(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
