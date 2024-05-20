package main

import (
	"InduksiTA/controllers"
	"InduksiTA/initializers"
	"InduksiTA/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
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
	category := router.Group("/category")
	userAuth.Use(middleware.RequiredAuth)

	router.Static("/images", "images/")

	userAuth.POST("/getUser", controllers.GetUserById)
	router.PUT("/user/foto/:id", controllers.UpdatePhotoProfile)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/resendOtp", controllers.ResendOtpEmailPassVer)

	category.GET("/machine", controllers.GetCategoryMachine)
	category.POST("/machine", controllers.CategoryMachine)

	router.GET("/store/items", controllers.GetStoreItems)
	router.POST("/store/items", controllers.StoreItems)

	router.POST("/sales", controllers.SalesReport)
	router.POST("/service", controllers.ServiceReport)
	router.GET("/sales", controllers.GetSalesReport)
	router.GET("/service", controllers.GetServiceReport)
	router.GET("/search", controllers.Search)

	viPay := router.Group("/viPay")

	viPay.POST("/profile", controllers.GetProfile)
	viPay.POST("/game-feature", controllers.GetGameOrder)
	viPay.POST("/list-game-price", controllers.ListGameHarga)
	viPay.POST("/get-nick-game", controllers.GetNickGame)
	viPay.POST("/top-up", controllers.TopUpGame)
	viPay.POST("/top-up-prepaid", controllers.TopUpPrepaid)
	viPay.POST("/list-prepaid-price", controllers.ListPrepaid)
	viPay.POST("/list-prepaid-order", controllers.GetPrepaidOrder)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	})

	komikCast := router.Group("/komikCast")

	komikCast.GET("/daftar-komik", func(c *gin.Context) {
		order := c.Query("order")
		page := c.Query("page")

		response, err := controllers.GetDaftarKomik(order, page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	})
	komikCast.GET("/fetch-data", controllers.GetDataHandler)
	komikCast.GET("/komik-info", controllers.GetKomikInfo)
	komikCast.GET("/search", controllers.SearchKomik)
	komikCast.GET("/genre", controllers.GetGenreInfo)
	komikCast.GET("/genre/komik", controllers.FetchComicsByGenre)

	router.Run()
}
