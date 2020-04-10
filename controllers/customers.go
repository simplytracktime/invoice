package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/simplytracktime/invoice/models"
)

// GET /customers

// FindCustomers returns all customers
func FindCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var customers []models.Customer
	db.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

//CreateCustomerInput used to create customers
type CreateCustomerInput struct {
	Name         string `json:"name" binding:"required"`
	EmailAddress string `json:"email_address" binding:"required"`
	Attention    string `json:"attention"`
	Address1     string `json:"address_1"`
	Address2     string `json:"address_2"`
}

// CreateCustomer a new customer
func CreateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create customer
	customer := models.Customer{
		Name:         input.Name,
		EmailAddress: input.EmailAddress,
		Attention:    input.Attention,
		Address1:     input.Address1,
		Address2:     input.Address2,
	}

	db.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// DeleteCustomer a customer
func DeleteCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var customer models.Customer
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
