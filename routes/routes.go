package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/login", login)
	r.POST("/register", register)

	r.GET("/products/:id", getProductsByCategory)

	r.POST("/payments", checkoutPayment)

	r.POST("/orders", checkoutOrder)

	r.GET("/cart-items/:id", getCartItemsByUser)
	r.POST("/cart-items/", createCartItem)
	r.DELETE("/cart-items/:id", deleteCartItem)
}
