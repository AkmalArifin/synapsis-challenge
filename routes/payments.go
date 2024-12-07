package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synapsis-challenge/models"
)

func getPayments(c *gin.Context) {
	payments, err := models.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func checkoutPayment(c *gin.Context) {
	var payment models.Payment
	err := c.ShouldBindJSON(&payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON"})
		return
	}

	err = payment.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payment)
}
