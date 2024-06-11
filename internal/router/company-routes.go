package router

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/bozhidarv/warehouse-manager/warehouse-manager-api/internal/db"
)

func AddCompanyRouter(rg *gin.RouterGroup) {
	companyRouter := rg.Group("/companies")
	companyRouter.GET("/", getCompanies)
	companyRouter.GET("/:id", getCompany)
	companyRouter.POST("/", createCompany)
	companyRouter.PUT("/:id", updateCompany)
	companyRouter.DELETE("/:id", deleteCompany)
}

func getCompanies(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	companies, err := dbConn.GetAllCompanies(c.Request.Context())
	if err != nil {
		c.Status(500)
		return
	}

	if companies == nil {
		companies = []db.Company{}
	}
	c.JSON(200, companies)
}

func getCompany(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	company, err := dbConn.GetCompanyById(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, company)
}

func createCompany(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	companyBody := db.CreateCompanyParams{}
	err = json.Unmarshal(bodyStr, &companyBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.CreateCompany(c.Request.Context(), companyBody)
	if err != nil {
		c.Status(500)
	}
	companyBody.ID = []byte(uuid.New().String())

	c.Status(201)
}

func updateCompany(c *gin.Context) {
	body := c.Request.Body
	defer body.Close()
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		c.Status(500)
	}

	companyBody := db.UpdateCompanyParams{}
	err = json.Unmarshal(bodyStr, &companyBody)
	if err != nil {
		c.Status(500)
	}

	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())

	dbConn := db.New(conn)
	err = dbConn.UpdateCompany(c.Request.Context(), companyBody)
	if err != nil {
		c.Status(500)
	}

	c.Status(201)
}

func deleteCompany(c *gin.Context) {
	conn := connectToDB(c)
	defer conn.Close(c.Request.Context())
	dbConn := db.New(conn)

	err := dbConn.DeleteCompany(c.Request.Context(), []byte(c.Param("id")))
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
