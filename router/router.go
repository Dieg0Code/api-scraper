package router

import (
	"net/http"

	"github.com/dieg0code/scraper-lab/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(productController *controller.ProductsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "API Running")
	})

	baseRouter := router.Group("/api/v1")
	productRouter := baseRouter.Group("/products")
	productRouter.GET("", productController.FindAll)
	productRouter.GET("/:productId", productController.FindById)
	productRouter.POST("", productController.Create)
	productRouter.POST("/update-data", productController.UpdateData)
	productRouter.PATCH("/:productId", productController.Update)
	productRouter.DELETE("/:productId", productController.Delete)

	return router
}
