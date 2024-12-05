package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synapsis-challenge/models"
)

func checkoutOrder(c *gin.Context) {
	var cart models.Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON"})
		return
	}

	order, orderItems, err := models.CreateOrderFromCart(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": order, "orderItems": orderItems})
}
