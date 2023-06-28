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
		publicRoutes.POST("/register-customer", controllers.RegisterCustomer)
		publicRoutes.POST("/register-tukang", controllers.RegisterTukang)
		publicRoutes.POST("/login", controllers.Login)
	}

	router.GET("/api/v1/checkIsAvailableEmail", controllers.CheckIsAvailableEmail)
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
				users.GET("/findEmail", controllers.FindUserByEmail)
				users.GET("/checkIsAvailableEmail", controllers.CheckIsAvailableEmail)
				users.GET("/findTukang", controllers.FindTukang)
				users.GET("/findTukang/:id", controllers.DetailTukang)
				// users.POST("/", controllers.CreateUser)
				users.PATCH("/:id", controllers.UpdateUser)
				users.DELETE("/:id", controllers.DeleteUser)
			}
			orders := v1.Group("/orders")
			{
				orders.POST("/:id/order", controllers.Order)
				orders.PUT("/accOrderByTukang/:id", controllers.AccOrderByTukang)
				orders.PUT("/rejectOrderByTukang/:id", controllers.RejectOrderByTukang)
				orders.GET("/statusOrderCustomerMenunggu", controllers.StatusOrderCustomerMenunggu)
				orders.GET("/statusOrderCustomerBerlangsung", controllers.StatusOrderCustomerBerlangsung)
				orders.GET("/statusOrderCustomerSelesai", controllers.StatusOrderCustomerSelesai)
				orders.GET("/statusOrderTukangMenunggu", controllers.StatusOrderTukangMenunggu)
				orders.GET("/statusOrderTukangBerlangsung", controllers.StatusOrderTukangBerlangsung)
				orders.GET("/statusOrderTukangSelesai", controllers.StatusOrderTukangSelesai)
				orders.PATCH("/ubahWaktu/:id", controllers.UbahWaktu)
			}
			categories := v1.Group("/categories")
			{
				categories.GET("/", controllers.GetAllCategories)
				categories.POST("/", controllers.CreateCategory)
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

/*curl -d '{"content":"A sample content"}' -H "Content-Type: application/json" -H "Authorization: Bearer <<JWT>>" -X POST http://localhost:8000/api/entry

curl -d '{"content":"A sample content"}' -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODI1ODk1MjYsImlhdCI6MTY4MjU4NzUyNiwiaWQiOjV9.Ygkg2HFm9HNuS42raazpTU179omt8OTjQSYv_2KOTkU" -X POST http://localhost:8000/api/entry

curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODI1ODk1MjYsImlhdCI6MTY4MjU4NzUyNiwiaWQiOjV9.Ygkg2HFm9HNuS42raazpTU179omt8OTjQSYv_2KOTkU" -X POST http://localhost:8000/api/v1/entry


curl -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2ODI1ODk1MjYsImlhdCI6MTY4MjU4NzUyNiwiaWQiOjV9.Ygkg2HFm9HNuS42raazpTU179omt8OTjQSYv_2KOTkU" -X GET http://localhost:8000/api/v1/entry
*/
