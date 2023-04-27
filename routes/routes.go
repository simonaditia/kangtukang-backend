package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/controllers"
	"github.com/simonaditia/kangtukang-backend/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	publicRoutes := router.Group("/auth")
	{
		publicRoutes.POST("/register", controllers.Register)
		publicRoutes.POST("/login", controllers.Login)
	}

	protectedRoutes := router.Group("/api")
	{
		protectedRoutes.Use(middleware.JWTAuthMiddleware())
		v1 := protectedRoutes.Group("/v1")
		{
			entries := v1.Group("/entry")
			{
				entries.POST("/", controllers.AddEntry)
				entries.GET("/", controllers.GetALlEntries)
			}
			users := v1.Group("/users")
			{
				users.GET("/", controllers.FindUsers)
				users.GET("/:id", controllers.FindUser)
				// users.POST("/", controllers.CreateUser)
				users.PATCH("/:id", controllers.UpdateUser)
				users.DELETE("/:id", controllers.DeleteUser)
			}
		}
	}

	/*v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.FindUsers)
			users.GET("/:id", controllers.FindUser)
			// users.POST("/", controllers.CreateUser)
			// users.POST("/", controllers.Register)
			users.PATCH("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}*/
	return router
}
