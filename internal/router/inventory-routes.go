package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddInventoryRouter(rg *gin.RouterGroup) {
	inventoryRouter := rg.Group("/inventory")
	inventoryRouter.GET("/", getInventorys)
	inventoryRouter.GET("/:id", getInventory)
	inventoryRouter.POST("/", createInventory)
	inventoryRouter.PUT("/:id", updateInventory)
	inventoryRouter.DELETE("/:id", deleteInventory)
}

func getInventorys(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	inventorys, err := dbConn.GetAllInventory(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if inventorys == nil {
		inventorys = []db.Inventory{}
	}
	c.JSON(200, inventorys)
}

func getInventory(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	inventory, err := dbConn.GetInventoryByMaterialId(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, inventory)
}

func createInventory(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	inventoryBody := db.CreateInventoryParams{}
	err = json.Unmarshal(bodyStr, &inventoryBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateInventory(c.Request.Context(), inventoryBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func updateInventory(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	inventoryBody := db.UpdateInventoryParams{}
	err = json.Unmarshal(bodyStr, &inventoryBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateInventory(c.Request.Context(), inventoryBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteInventory(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteInventory(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
