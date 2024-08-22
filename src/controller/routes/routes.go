package routes

import (
	controller "github.com/enioaires/crud-go/src/controller/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	// Get user by id
	r.GET("/user/:id", controller.FindUserById)

	// Get user by email
	r.GET("/user/email/:email", controller.FindUserByEmail)

	// Insert user
	r.POST("/user", controller.CreateUser)

	// Update user
	r.PUT("/user/:id", controller.UpdateUser)

	// Delete user
	r.DELETE("/user/:id", controller.DeleteUser)
}
