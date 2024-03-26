package service

import (
	"github.com/dieg0code/scraper-lab/data/request"
	"github.com/dieg0code/scraper-lab/data/response"
)

type ProductsService interface {
	Create(product request.CreateProductsRequest)
	Update(product request.UpdateProductsRequest)
	UpdateData(updateData request.UpdateDataRequest) bool
	Delete(productId int)
	FindByID(productId int) response.ProductsResponse
	FindAll() []response.ProductsResponse
}
