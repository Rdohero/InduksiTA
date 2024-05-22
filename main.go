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
	category.GET("/spare/part", controllers.GetCategorySparePart)
	category.POST("/spare/part", controllers.CategorySparePart)

	router.GET("/store/items", controllers.GetStoreItems)
	router.POST("/store/items", controllers.StoreItems)

	router.GET("/spare/part", controllers.GetSparePart)
	router.POST("/spare/part", controllers.SparePart)

	router.POST("/sales", controllers.SalesReport)
	router.POST("/service", controllers.ServiceReport)
	router.GET("/sales", controllers.GetSalesReport)
	router.GET("/service", controllers.GetServiceReport)

	search := router.Group("search")

	search.GET("/machine", controllers.SearchMachine)
	search.GET("/sparePart", controllers.SearchSparePart)

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
	komikCast.GET("/daftar-komik/eleceed", func(c *gin.Context) {
		response := gin.H{
			"chapters": []gin.H{
				{
					"chapter": "298",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-298-bahasa-indonesia/",
					"time":    "1 week ago",
				},
				{
					"chapter": "297",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-297-bahasa-indonesia/",
					"time":    "2 weeks ago",
				},
				{
					"chapter": "296",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-296-bahasa-indonesia/",
					"time":    "2 weeks ago",
				},
				{
					"chapter": "295",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-295-bahasa-indonesia/",
					"time":    "3 weeks ago",
				},
				{
					"chapter": "294",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-294-bahasa-indonesia/",
					"time":    "4 weeks ago",
				},
				{
					"chapter": "293",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-293-bahasa-indonesia/",
					"time":    "1 month ago",
				},
				{
					"chapter": "292",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-292-bahasa-indonesia/",
					"time":    "1 month ago",
				},
				{
					"chapter": "291",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-291-bahasa-indonesia/",
					"time":    "2 months ago",
				},
				{
					"chapter": "290",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-290-bahasa-indonesia/",
					"time":    "2 months ago",
				},
				{
					"chapter": "289",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-289-bahasa-indonesia/",
					"time":    "2 months ago",
				},
				{
					"chapter": "288",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-288-bahasa-indonesia/",
					"time":    "2 months ago",
				},
				{
					"chapter": "287",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-287-bahasa-indonesia/",
					"time":    "3 months ago",
				},
				{
					"chapter": "286",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-286-bahasa-indonesia/",
					"time":    "3 months ago",
				},
				{
					"chapter": "285",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-285-bahasa-indonesia/",
					"time":    "3 months ago",
				},
				{
					"chapter": "284",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-284-bahasa-indonesia/",
					"time":    "3 months ago",
				},
				{
					"chapter": "283",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-283-bahasa-indonesia/",
					"time":    "4 months ago",
				},
				{
					"chapter": "282",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-282-bahasa-indonesia/",
					"time":    "4 months ago",
				},
				{
					"chapter": "281",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-281-bahasa-indonesia/",
					"time":    "4 months ago",
				},
				{
					"chapter": "280",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-280-bahasa-indonesia/",
					"time":    "4 months ago",
				},
				{
					"chapter": "279",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-279-bahasa-indonesia/",
					"time":    "4 months ago",
				},
				{
					"chapter": "278",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-278-bahasa-indonesia/",
					"time":    "5 months ago",
				},
				{
					"chapter": "277",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-277-bahasa-indonesia/",
					"time":    "5 months ago",
				},
				{
					"chapter": "276",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-276-bahasa-indonesia/",
					"time":    "5 months ago",
				},
				{
					"chapter": "275",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-275-bahasa-indonesia/",
					"time":    "5 months ago",
				},
				{
					"chapter": "274",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-274-bahasa-indonesia/",
					"time":    "6 months ago",
				},
				{
					"chapter": "273",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-273-bahasa-indonesia/",
					"time":    "6 months ago",
				},
				{
					"chapter": "272",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-272-bahasa-indonesia/",
					"time":    "6 months ago",
				},
				{
					"chapter": "271.5",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-271-5-bahasa-indonesia/",
					"time":    "6 months ago",
				},
				{
					"chapter": "271",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-271-bahasa-indonesia/",
					"time":    "7 months ago",
				},
				{
					"chapter": "270 Fix",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-270-bahasa-indonesia/",
					"time":    "7 months ago",
				},
				{
					"chapter": "269",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-269-bahasa-indonesia/",
					"time":    "7 months ago",
				},
				{
					"chapter": "268",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-268-bahasa-indonesia/",
					"time":    "7 months ago",
				},
				{
					"chapter": "267",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-267-bahasa-indonesia/",
					"time":    "7 months ago",
				},
				{
					"chapter": "266",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-266-bahasa-indonesia/",
					"time":    "8 months ago",
				},
				{
					"chapter": "265",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-265-bahasa-indonesia/",
					"time":    "8 months ago",
				},
				{
					"chapter": "264",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-264-bahasa-indonesia/",
					"time":    "8 months ago",
				},
				{
					"chapter": "263",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-263-bahasa-indonesia/",
					"time":    "8 months ago",
				},
				{
					"chapter": "262",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-262-bahasa-indonesia/",
					"time":    "9 months ago",
				},
				{
					"chapter": "261",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-261-bahasa-indonesia/",
					"time":    "9 months ago",
				},
				{
					"chapter": "260",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-260-bahasa-indonesia/",
					"time":    "9 months ago",
				},
				{
					"chapter": "259",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-259-bahasa-indonesia/",
					"time":    "9 months ago",
				},
				{
					"chapter": "258",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-258-bahasa-indonesia/",
					"time":    "10 months ago",
				},
				{
					"chapter": "257",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-257-bahasa-indonesia/",
					"time":    "10 months ago",
				},
				{
					"chapter": "256",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-256-bahasa-indonesia/",
					"time":    "10 months ago",
				},
				{
					"chapter": "255",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-255-bahasa-indonesia/",
					"time":    "10 months ago",
				},
				{
					"chapter": "254",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-254-bahasa-indonesia/",
					"time":    "11 months ago",
				},
				{
					"chapter": "253",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-253-bahasa-indonesia/",
					"time":    "11 months ago",
				},
				{
					"chapter": "252",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-252-bahasa-indonesia/",
					"time":    "11 months ago",
				},
				{
					"chapter": "251",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-251-bahasa-indonesia/",
					"time":    "11 months ago",
				},
				{
					"chapter": "250",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-250-bahasa-indonesia/",
					"time":    "11 months ago",
				},
				{
					"chapter": "249",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-249-bahasa-indonesia/",
					"time":    "12 months ago",
				},
				{
					"chapter": "248",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-248-bahasa-indonesia/",
					"time":    "12 months ago",
				},
				{
					"chapter": "247",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-247-bahasa-indonesia/",
					"time":    "12 months ago",
				},
				{
					"chapter": "246",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-246-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "245",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-245-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "244",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-244-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "243",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-243-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "242",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-242-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "241",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-241-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "240",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-240-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "239",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-239-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "238",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-238-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "237",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-237-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "236",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-236-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "235",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-235-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "234",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-234-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "233",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-233-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "232",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-232-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "231",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-231-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "230",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-230-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "229",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-229-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "228",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-228-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "227",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-227-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "226",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-226-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "225",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-225-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "224",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-224-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "223",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-223-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "222",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-222-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "221",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-221-bahasa-indonesia/",
					"time":    "1 year ago",
				},
				{
					"chapter": "220",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-220-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "219",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-219-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "218",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-218-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "217",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-217-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "216",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-216-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "215",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-215-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "214",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-214-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "213",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-213-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "212",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-212-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "211",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-211-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "210",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-210-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "209",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-209-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "208",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-208-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "207",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-207-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "206",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-206-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "205",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-205-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "204",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-204-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "203",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-203-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "202",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-202-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "201",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-201-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "200",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-200-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "199",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-199-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "198",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-198-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "197",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-197-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "196",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-196-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "195",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-195-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "194",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-194-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "193",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-193-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "192",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-192-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "191",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-191-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "190",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-190-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "189",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-189-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "188",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-188-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "187",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-187-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "186",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-186-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "185",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-185-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "184",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-184-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "183",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-183-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "182",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-182-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "181",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-181-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "180",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-180-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "179",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-179-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "178",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-178-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "177",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-177-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "176",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-176-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "175",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-175-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "174",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-174-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "173",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-173-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "172",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-172-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "171",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-171-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "170",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-170-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "169",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-169-bahasa-indonesia/",
					"time":    "2 years ago",
				},
				{
					"chapter": "168",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-168-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "167",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-167-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "166",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-166-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "165",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-165-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "164",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-164-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "163",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-163-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "162",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-162-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "161",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-161-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "160",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-160-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "159",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-159-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "158",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-158-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "157",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-157-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "156",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-156-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "155",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-155-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "154",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-154-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "153",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-153-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "152",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-152-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "151",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-151-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "150",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-150-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "149",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-149-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "148",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-148-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "147",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-147-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "146",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-146-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "145",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-145-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "144",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-144-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "143",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-143-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "142",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-142-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "141",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-141-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "140",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-140-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "139",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-139-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "138",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-138-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "137",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-137-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "136",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-136-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "135",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-135-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "134",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-134-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "133",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-133-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "132",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-132-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "131",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-131-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "130",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-130-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "129",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-129-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "128",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-128-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "127",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-127-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "126",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-126-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "125",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-125-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "124",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-124-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "123",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-123-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "122",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-122-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "121",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-121-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "120",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-120-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "119",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-119-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "118",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-118-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "117",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-117-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "116",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-116-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "115",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-115-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "114",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-114-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "113",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-113-bahasa-indonesia/",
					"time":    "3 years ago",
				},
				{
					"chapter": "112",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-112-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "111",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-111-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "110",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-110-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "109",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-109-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "108",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-108-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "107",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-107-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "106",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-106-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "105",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-105-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "104",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-104-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "103",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-103-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "102",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-102-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "101",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-101-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "100",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-100-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "99",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-99-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "98",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-98-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "97",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-97-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "96",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-96-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "95",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-95-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "94",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-94-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "93",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-93-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "92",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-92-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "91",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-91-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "90",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-90-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "89",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-89-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "88",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-88-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "87",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-87-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "86",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-86-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "85",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-85-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "84",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-84-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "83",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-83-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "82",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-82-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "81",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-81-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "80",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-80-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "79",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-79-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "78",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-78-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "77",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-77-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "76",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-76-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "75",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-75-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "74",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-74-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "73",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-73-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "72",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-72-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "71",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-71-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "70",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-70-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "69",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-69-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "68",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-68-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "67",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-67-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "66",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-66-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "65",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-65-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "64",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-64-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "63",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-63-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "62",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-62-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "61",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-61-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "60",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-60-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "59",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-59-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "58",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-58-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "57",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-57-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "56",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-56-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "55",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-55-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "54",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-54-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "53",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-53-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "52",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-52-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "51",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-51-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "50",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-50-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "49",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-49-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "48",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-48-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "47",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-47-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "46",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-46-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "45",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-45-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "44",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-44-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "43",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-43-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "42",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-42-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "41",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-41-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "40",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-40-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "39",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-39-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "38",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-38-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "37",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-37-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "36",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-36-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "35",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-35-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "34",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-34-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "33",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-33-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "32",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-32-bahasa-indonesia/",
					"time":    "4 years ago",
				},
				{
					"chapter": "31",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-31-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "30",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-30-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "29",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-29-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "28",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-28-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "27",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-27-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "26",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-26-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "25",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-25-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "24",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-24-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "23",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-23-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "22",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-22-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "21",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-21-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "20",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-20-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "19",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-19-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "18",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-18-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "17",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-17-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "16",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-16-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "15",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-15-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "14",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-14-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "13",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-13-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "12",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-12-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "11",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-11-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "10",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-10-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "09",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-09-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "08",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-08-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "07",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-07-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "06",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-06-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "05",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-05-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "04",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-04-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "03",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-03-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "02",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-02-bahasa-indonesia/",
					"time":    "5 years ago",
				},
				{
					"chapter": "01",
					"link":    "https://komikcast.lol/chapter/eleceed-chapter-01-bahasa-indonesia/",
					"time":    "5 years ago",
				},
			},
			"genre": []gin.H{
				{
					"Name": "Ongoing",
					"Link": "",
				},
				{
					"Name": "Manhwa",
					"Link": "https://komikcast.lol/type/manhwa/",
				},
				{
					"Name": "Action",
					"Link": "https://komikcast.lol/genres/action/",
				},
				{
					"Name": "Fantasy",
					"Link": "https://komikcast.lol/genres/fantasy/",
				},
				{
					"Name": "Supernatural",
					"Link": "https://komikcast.lol/genres/supernatural/",
				},
			},
			"link":     "https://komikcast.lol/komik/eleceed/",
			"sinopsis": "Kaiden – Pengguna kemampuan misterius yang bersembunyi di dalam tubuh kucing jalanan. Dia kemudian dijemput oleh Jiwoo setelah terluka setelah berkelahi dengan pengguna kemampuan lain. Ia memiliki kepribadian yang keras kepala dan suka memerintah. Jiwoo – anak SMA yang energik dan banyak bicara yang suka kucing. Ia sangat baik dan juga tampaknya memiliki kemampuan khusus.",
			"title":    "Eleceed Bahasa Indonesia",
		}

		c.JSON(http.StatusOK, response)
	})

	router.Run()
}
