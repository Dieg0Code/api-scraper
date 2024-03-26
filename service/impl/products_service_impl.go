package impl

import (
	"fmt"

	"github.com/dieg0code/scraper-lab/data/request"
	"github.com/dieg0code/scraper-lab/data/response"
	"github.com/dieg0code/scraper-lab/helper"
	"github.com/dieg0code/scraper-lab/model"
	"github.com/dieg0code/scraper-lab/repository"
	"github.com/dieg0code/scraper-lab/service"
	"github.com/go-playground/validator/v10"
	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog/log"
)

type CategoryInfo struct {
	Category string
	MaxPage  int
}

type ProductsServiceImpl struct {
	ProductRepository repository.ProductsRepository
	Validate          *validator.Validate
}

func NewProductsServiceImpl(productRepository repository.ProductsRepository, validate *validator.Validate) service.ProductsService {
	return &ProductsServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

// Create implements service.ProductsService.
func (p *ProductsServiceImpl) Create(product request.CreateProductsRequest) {
	err := p.Validate.Struct(product)
	helper.ErrorPanic(err)
	productModel := model.Product{
		Name:          product.Name,
		Category:      product.Category,
		OriginalPrice: product.OriginalPrice,
		DiscountPrice: product.DiscountPrice,
		Supermarket:   product.Supermarket,
	}
	p.ProductRepository.SaveProduct(productModel)
}

// Delete implements service.ProductsService.
func (p *ProductsServiceImpl) Delete(productId int) {
	p.ProductRepository.Delete(productId)
}

// FindAll implements service.ProductsService.
func (p *ProductsServiceImpl) FindAll() []response.ProductsResponse {
	result := p.ProductRepository.FindAll()

	var products []response.ProductsResponse
	for _, product := range result {
		product := response.ProductsResponse{
			Id:            product.Id,
			Name:          product.Name,
			Category:      product.Category,
			OriginalPrice: product.OriginalPrice,
			DiscountPrice: product.DiscountPrice,
			Supermarket:   product.Supermarket,
		}
		products = append(products, product)
	}

	return products
}

// FindByID implements service.ProductsService.
func (p *ProductsServiceImpl) FindByID(productId int) response.ProductsResponse {
	productData, err := p.ProductRepository.FindByID(productId)
	helper.ErrorPanic(err)

	productResponse := response.ProductsResponse{
		Id:            productData.Id,
		Name:          productData.Name,
		Category:      productData.Category,
		OriginalPrice: productData.OriginalPrice,
		DiscountPrice: productData.DiscountPrice,
		Supermarket:   productData.Supermarket,
	}

	return productResponse
}

// Update implements service.ProductsService.
func (p *ProductsServiceImpl) Update(product request.UpdateProductsRequest) {
	productData, err := p.ProductRepository.FindByID(product.Id)
	helper.ErrorPanic(err)
	productData.Name = product.Name
	p.ProductRepository.Update(productData)
}

// UpdateData implements service.ProductsService.
func (p *ProductsServiceImpl) UpdateData(updateData request.UpdateDataRequest) bool {
	if updateData.Update {
		// Clear all products
		p.ProductRepository.ClearProducts()

		categories := []CategoryInfo{
			{Category: "bebidas-alcoholicas", MaxPage: 10},
			{Category: "bebidas-jugos-y-aguas", MaxPage: 8},
			{Category: "carniceria", MaxPage: 2},
			{Category: "cuidado-personal", MaxPage: 8},
			{Category: "desayuno", MaxPage: 5},
			{Category: "despensa", MaxPage: 13},
			{Category: "dulces-y-snacks", MaxPage: 5},
			{Category: "ferreteria", MaxPage: 1},
			{Category: "la-gran-feria-cugat", MaxPage: 3},
			{Category: "del-mundo-a-tu-despensa", MaxPage: 7},
			{Category: "lacteos", MaxPage: 7},
			{Category: "limpieza-y-aseo", MaxPage: 15},
			{Category: "mascotas", MaxPage: 1},
			{Category: "mundo-bebe", MaxPage: 3},
			{Category: "mundo-congelados", MaxPage: 9},
			{Category: "navidad", MaxPage: 1},
			{Category: "panaderia-y-pasteleria", MaxPage: 2},
			{Category: "preparados", MaxPage: 1},
			{Category: "quesos-y-fiambreria", MaxPage: 7},
		}

		for _, categoryInfo := range categories {
			products := scrape("cugat.cl/categoria-producto", categoryInfo.MaxPage, categoryInfo.Category)
			for _, product := range products {
				productModel := model.Product{
					Name:          product.Name,
					Category:      product.Category,
					OriginalPrice: product.OriginalPrice,
					DiscountPrice: product.DiscountPrice,
					Supermarket:   product.Supermarket,
				}
				p.ProductRepository.SaveProduct(productModel)
				log.Info().Msg("Product updated")
			}
		}
	}
	return true
}

func scrape(baseURL string, maxPage int, category string) []model.Product {
	c := colly.NewCollector()

	var products []model.Product

	c.OnHTML(".product-small.box", func(e *colly.HTMLElement) {
		name := e.ChildText(".name.product-title a")
		category := e.ChildText(".category")
		originalPrice := e.ChildText(".price del .woocommerce-Price-amount.amount")
		discountPrice := e.ChildText(".price ins .woocommerce-Price-amount.amount")
		supermarket := "Cugat"

		if discountPrice == "" {
			originalPrice = e.ChildText(".price .woocommerce-Price-amount.amount")
		}

		products = append(products, model.Product{
			Name:          name,
			Category:      category,
			OriginalPrice: originalPrice,
			DiscountPrice: discountPrice,
			Supermarket:   supermarket,
		})
	})

	for i := 1; i <= maxPage; i++ {
		c.Visit(fmt.Sprintf("https://%s/%s/page/%d/", baseURL, category, i))
		log.Info().Msgf("Scraping page https://%s/%s/page/%d/", baseURL, category, i)
	}

	return products
}
