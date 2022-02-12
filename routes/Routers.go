package routes

import (
	"github.com/gin-gonic/gin"
	"mvc/controllers"
	"mvc/middlewares"
)

func SetRouter() *gin.Engine  {
	router := gin.Default()
	router.TrustedPlatform = "True-Client-IP"
	routes := router.Group("api")
	{
		routes.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		userGroup := routes.Group("user", middlewares.VerifyToken())
		{
			userGroup.POST("", controllers.CreateUser)
			userGroup.GET("", controllers.GetUserList)
			userGroup.GET("/:id", controllers.GetUserDetail)
			userGroup.PATCH("/:id", controllers.UpdateUser)
			userGroup.DELETE("/:id", controllers.DeleteUser)
		}
	}

	return router
}
