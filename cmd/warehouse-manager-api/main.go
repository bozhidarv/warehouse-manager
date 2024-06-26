package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/router"
)

func main() {
	gin.ForceConsoleColor()

	r := gin.Default()

	router.AddMaterialRouter(&r.RouterGroup)
	router.AddSupplierRouter(&r.RouterGroup)
	router.AddCompanyRouter(&r.RouterGroup)
	router.AddUnitRouter(&r.RouterGroup)
	router.AddUserRouter(&r.RouterGroup)
	router.AddRecipeRouter(&r.RouterGroup)
	router.AddInventoryRouter(&r.RouterGroup)
	router.AddTransactionHistoryRouter(&r.RouterGroup)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	log.Fatal(r.Run(":8080"))
}
