package repository

import "github.com/dieg0code/scraper-lab/model"

type ProductsRepository interface {
	SaveProduct(product model.Product)
	Update(product model.Product)
	Delete(productId int)
	ClearProducts()
	FindByID(productId int) (product model.Product, err error)
	FindAll() []model.Product
}
