package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func connectToDB(c *gin.Context) *pgx.Conn {
	connStr := "postgres://postgres:warehouse-manager@localhost:5432/postgres?sslmode=disable"
	conn, err := pgx.Connect(c.Request.Context(), connStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error connecting to the database",
		})
		return nil
	}
	return conn
}
