package main

import (
	"github.com/gin-gonic/gin"
	"github.com/simplytracktime/invoice/controllers"
	"github.com/simplytracktime/invoice/models"
)

func main() {

	r := gin.Default()

	db := models.SetupModels() //new

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/customers", controllers.FindCustomers)
	r.POST("/customers", controllers.CreateCustomer)

	r.DELETE("/customers/:id", controllers.DeleteCustomer)

	r.Run()
}
