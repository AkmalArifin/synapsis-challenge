package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/synapsis-challenge/models"
)

func createCartItem(c *gin.Context) {
	var cartItem models.CartItem
	err := c.ShouldBindJSON(&cartItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON"})
		return
	}

	// TODO: If exist, add one
	err = cartItem.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data"})
		return
	}

	c.JSON(http.StatusOK, cartItem)
}

func getCartItemsByUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	cart, err := models.GetCartByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data", "cart": cart})
		return
	}

	cartItems, err := models.GetCartItemsByCart(cart.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func deleteCartItem(c *gin.Context) {
	cartID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	cartItem, err := models.GetCartItemByID(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	err = cartItem.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been deleted"})
}
