package routes

import (
	"github.com/gin-contrib/cors"
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

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "halo halo")
	})

	router.Use(cors.Default())

	publicRoutes := router.Group("/auth")
	{
		publicRoutes.POST("/register-customer", controllers.RegisterCustomer)
		publicRoutes.POST("/register-tukang", controllers.RegisterTukang)
		publicRoutes.POST("/login", controllers.Login)
		// publicRoutes.POST("/v2/register-customer", controllers.RegisterCustomerV2)
		publicRoutes.POST("/v2/register-customer", controllers.RegisterCustomerV2)
		publicRoutes.POST("/v2/register-tukang", controllers.RegisterTukangV2)
		publicRoutes.POST("/v2/login", controllers.LoginV2)
	}

	router.GET("/api/v1/checkIsAvailableEmail", controllers.CheckIsAvailableEmail)
	router.GET("/api/v1/checkIsAvailableNoTelp", controllers.CheckIsAvailableNoTelp)
	router.GET("/website/api/v1/users/allCustomer", controllers.FindAllCustomer)
	router.GET("/website/api/v1/users/allTukang", controllers.FindAllTukang)
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
				users.GET("/allCustomer", controllers.FindAllCustomer)
				users.GET("/allTukang", controllers.FindAllTukang)
				users.GET("/findEmail", controllers.FindUserByEmail)
				users.GET("/checkIsAvailableEmail", controllers.CheckIsAvailableEmail)
				users.GET("/findTukang", controllers.FindTukang)
				users.GET("/findTukang/:id", controllers.DetailTukang)
				// users.POST("/", controllers.CreateUser)
				users.PATCH("/:id", controllers.UpdateUser)
				users.DELETE("/:id", controllers.DeleteUser)
				users.POST("/:id/categories/:categoryID", controllers.AddUserCategory)
				users.PUT("/:id/categories", controllers.UpdateUserCategories)
			}
			orders := v1.Group("/orders")
			{
				orders.POST("/:id/order", controllers.Order)
				orders.PUT("/cancelOrderByCustomer/:id", controllers.CancelOrderByCustomer)
				orders.PUT("/accOrderByTukang/:id", controllers.AcceptOrderByTukang)
				orders.PUT("/rejectOrderByTukang/:id", controllers.RejectOrderByTukang)
				orders.PUT("/doneOrderByTukang/:id", controllers.DoneOrderByTukang)
				orders.GET("/readOrderByTukang", controllers.ReadOrderByTukang)
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
				categories.GET("/:id", controllers.GetCategory)
				categories.POST("/:id/users/:userID", controllers.AddCategoryUser)
				// categories.GET("/", controllers.GetAllCategories)
				// categories.POST("/", controllers.CreateCategory)
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
