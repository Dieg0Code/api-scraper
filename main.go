package main

import (
	"net/http"

	"github.com/dieg0code/scraper-lab/config"
	"github.com/dieg0code/scraper-lab/controller"
	"github.com/dieg0code/scraper-lab/helper"
	"github.com/dieg0code/scraper-lab/model"
	repo "github.com/dieg0code/scraper-lab/repository/impl"
	"github.com/dieg0code/scraper-lab/router"
	service "github.com/dieg0code/scraper-lab/service/impl"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("[Main] :: Server Running")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("products").AutoMigrate(&model.Product{})

	// Repository
	productsRepository := repo.NewProductsRepositoryImpl(db)

	// Service
	productService := service.NewProductsServiceImpl(productsRepository, validate)

	// Controller
	productController := controller.NewProductsController(productService)

	// Router
	routes := router.NewRouter(productController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
