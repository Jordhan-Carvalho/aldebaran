package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordhan-carvalho/aldebaran/internal/controllers"
	"github.com/jordhan-carvalho/aldebaran/internal/services"
)

func InitWebServer() {
	router := gin.Default()
	r := WebRouter(router)

	r.Run()
}

func WebRouter(r *gin.Engine) (router *gin.Engine) {

	// Health check

	/* // Routes - Status Check
	status := controllers.NewStatusController(svcInfo, dbMgr)
	router.GET("/status", status.CheckStatus) // /status

	// Dependencies for controllers
	d, _ := dbMgr.Database()
	orders := db.NewOrderDataService(d)

	// Routes - Seed DB
	if util.IsDevMode(svcInfo.Environment) {
		seed := controllers.NewSeedController(orders)
		router.POST("/seedDB", seed.SeedDB) // /seedDB
	}

	// Routes - Orders
	v1 := router.Group("/api/v1")
	{
		ordersGroup := v1.Group("orders")
		{
			orders := controllers.NewOrdersController(orders)
			ordersGroup.GET("", orders.GetAll)            // api/v1/orders
			ordersGroup.GET("/:id", orders.GetById)       // api/v1/orders/:id
			ordersGroup.POST("", orders.Post)             // api/v1/orders
			ordersGroup.PUT("", orders.Post)              // api/v1/orders
			ordersGroup.DELETE("/:id", orders.DeleteById) // api/v1/orders/:id
		}
	} */
	r.GET("/health", func(ctx *gin.Context) {
		statusController := controllers.StatusController{}
		statusService := services.StatusService{}
		message := statusController.GetAldebaranStatus(&statusService)

		ctx.JSON(http.StatusOK, message)
	})

	// router.Run()
	router = r
	return
}
