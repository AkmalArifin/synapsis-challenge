package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/synapsis-challenge/models"
)

func getProductsByCategory(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	products, err := models.GetProductsByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, products)
}
