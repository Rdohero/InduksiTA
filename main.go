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
	userAuth.Use(middleware.RequiredAuth)

	userAuth.POST("/getUser", controllers.GetUserById)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/resendOtp", controllers.ResendOtpEmailPassVer)
	router.POST("/emailve", controllers.OtpEmailVer)

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

	router.Run()
}
