package main

import (
	"InduksiTA/controllers"
	"InduksiTA/initializers"
	"InduksiTA/middleware"
	"InduksiTA/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.MigrateDatabase()
}

func main() {
	router := gin.Default()
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "ngrok-skip-browser-warning", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))

	userAuth := router.Group("/userAuth")
	userAuth.Use(middleware.RequiredAuth)

	router.Static("/images", "images/")

	router.GET("/user/all", controllers.GetAllUser)
	router.GET("/role", controllers.GetRole)

	userAuth.POST("/getUser", controllers.GetUserById)
	router.PUT("/user/foto/:id", controllers.UpdatePhotoProfile)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/login/owner", controllers.LoginOwner)
	router.PUT("/user", controllers.ChangeProfileUser)
	router.DELETE("/user/:id", controllers.DeletedUser)
	router.POST("/forgot/password", controllers.ForgotPassword)
	router.POST("/otp", controllers.ResendOtpEmailPassVer)

	router.GET("/category", controllers.GetCategory)
	router.POST("/category", controllers.CategoryPost)
	router.PUT("/category", controllers.EditCategory)
	router.DELETE("/category/:id", controllers.DeletedCategory)

	router.GET("/store/items", controllers.GetStoreItems)
	router.POST("/store/items", controllers.StoreItems)
	router.PUT("/store/items", controllers.EditStoreItems)
	router.DELETE("/store/items/:id", controllers.DeleteStoreItems)

	router.GET("/spare/part", controllers.GetSparePart)
	router.POST("/spare/part", controllers.SparePart)
	router.PUT("/spare/part", controllers.EditSparePart)
	router.DELETE("/spare/part/:id", controllers.DeleteSparePart)

	router.POST("/sales", controllers.SalesReport)
	router.POST("/service", controllers.ServiceReport)
	router.GET("/sales", controllers.GetSalesReport)
	router.GET("/sales/day/last", controllers.GetSalesReportsLastDays)
	router.GET("/sales/day/range", controllers.GetSalesReportsByDateRange)
	router.GET("/service", controllers.GetServiceReport)
	router.GET("/service/:id", controllers.GetServiceReportByUserID)
	router.GET("/service/status/:id", controllers.GetServiceReportByStatusID)
	router.PUT("/service", controllers.EditServiceReport)
	router.GET("/service/day/last", controllers.GetServiceReportsLastDays)
	router.GET("/service/day/range", controllers.GetServiceReportsByDateRange)
	router.DELETE("/sales/:id", controllers.DeletedSalesReport)

	router.GET("/status", controllers.GetStatus)

	search := router.Group("search")
	preOrder := router.Group("preOrder")

	search.GET("/machine", controllers.SearchMachine)
	search.GET("/sparePart", controllers.SearchSparePart)
	search.GET("/sales", controllers.SearchSales)
	search.GET("/service", controllers.SearchService)
	preOrder.POST("/store/items", controllers.PreOrderStoreItems)
	preOrder.POST("/spare/part", controllers.PreOrderSparePart)

	router.GET("/update-stock", func(c *gin.Context) {
		stockStr := c.Query("stock")

		stock, err := strconv.Atoi(stockStr)
		if err != nil || stock < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock value"})
			return
		}

		var storeItems []models.StoreItems
		var spareParts []models.SparePart

		initializers.DB.Find(&storeItems)
		for _, item := range storeItems {
			if item.Quantity == 0 {
				initializers.DB.Model(&item).Update("quantity", stock)
			}
		}

		initializers.DB.Find(&spareParts)
		for _, part := range spareParts {
			if part.Quantity == 0 {
				initializers.DB.Model(&part).Update("quantity", stock)
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Stock updated successfully"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	})

	router.Run()
}
