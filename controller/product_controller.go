package controller

import (
	"net/http"
	"strconv"

	"github.com/dieg0code/scraper-lab/data/request"
	"github.com/dieg0code/scraper-lab/data/response"
	"github.com/dieg0code/scraper-lab/helper"
	"github.com/dieg0code/scraper-lab/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ProductsController struct {
	productService service.ProductsService
}

func NewProductsController(service service.ProductsService) *ProductsController {
	return &ProductsController{
		productService: service,
	}
}

func (controller *ProductsController) Create(ctx *gin.Context) {
	log.Info().Msg("[ProductsController] :: Create Product")
	createProductRequest := request.CreateProductsRequest{}
	err := ctx.ShouldBindJSON(&createProductRequest)
	helper.ErrorPanic(err)

	controller.productService.Create(createProductRequest)
	webResponse := response.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) Update(ctx *gin.Context) {
	log.Info().Msg("[ProductsController] :: Update Product")
	updateProductsResquest := request.UpdateProductsRequest{}
	err := ctx.ShouldBindJSON(&updateProductsResquest)
	helper.ErrorPanic(err)

	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	updateProductsResquest.Id = id

	controller.productService.Update(updateProductsResquest)

	webResponse := response.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) Delete(ctx *gin.Context) {
	log.Info().Msg("[ProductsController] :: Delete Product")
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	controller.productService.Delete(id)

	webResponse := response.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) FindById(ctx *gin.Context) {
	log.Info().Msg("[ProductsController] :: Find Product By Id")
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)

	productResponse := controller.productService.FindByID(id)

	if productResponse.Id == 0 {
		webResponse := response.BaseResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   productResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("[ProductsController] :: Find All Product")
	productsResponse := controller.productService.FindAll()
	webResponse := response.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   productsResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductsController) UpdateData(ctx *gin.Context) {
	log.Info().Msg("[ProductsController] :: Update Data")
	updateDataRequest := request.UpdateDataRequest{}
	err := ctx.ShouldBindJSON(&updateDataRequest)
	helper.ErrorPanic(err)

	isSuccess := controller.productService.UpdateData(updateDataRequest)
	if !isSuccess {
		webResponse := response.BaseResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := response.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
