package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synapsis-challenge/models"
)

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorized account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login success"})
}

func register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON"})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not store data"})
		return
	}

	_, err = models.CreateCartByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not store data"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
