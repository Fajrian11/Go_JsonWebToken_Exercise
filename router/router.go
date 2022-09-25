package router

import (
	"Go_JsonWebToken_Exercise/controllers"
	"Go_JsonWebToken_Exercise/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := router.Group("products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateController)
	}

	return router
}
