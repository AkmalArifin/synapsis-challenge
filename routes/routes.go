package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/users", getUsers)
	r.POST("/login", login)
	r.POST("/register", register)

	r.GET("/products", getProducts)
	r.GET("/products/:id", getProductsByCategory)

	r.GET("/payments", getPayments)
	r.POST("/payments", checkoutPayment)

	r.GET("/orders", getOrders)
	r.GET("/order-items", getOrderItems)
	r.POST("/orders", checkoutOrder)

	r.GET("/carts", getCarts)
	r.GET("/cart-items/:id", getCartItemsByUser)
	r.POST("/cart-items/", createCartItem)
	r.DELETE("/cart-items/:id", deleteCartItem)
}
