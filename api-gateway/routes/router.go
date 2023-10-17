package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Get_gin_cors_func(), middleware.InitMiddleware(service))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.Any("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)
	}

	return ginRouter
}
