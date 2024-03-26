package impl

import (
	"errors"

	"github.com/dieg0code/scraper-lab/data/request"
	"github.com/dieg0code/scraper-lab/helper"
	"github.com/dieg0code/scraper-lab/model"
	"github.com/dieg0code/scraper-lab/repository"
	"gorm.io/gorm"
)

type ProductsRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductsRepositoryImpl(Db *gorm.DB) repository.ProductsRepository {
	return &ProductsRepositoryImpl{Db: Db}
}

// Delete implements repository.ProductsRepository.
func (p *ProductsRepositoryImpl) Delete(productId int) {
	var product model.Product
	result := p.Db.Where("id = ?", productId).Delete(&product)
	helper.ErrorPanic(result.Error)
}

// FindAll implements repository.ProductsRepository.
func (p *ProductsRepositoryImpl) FindAll() []model.Product {
	var products []model.Product
	result := p.Db.Find(&products)
	helper.ErrorPanic(result.Error)
	return products
}

// FindByID implements repository.ProductsRepository.
func (p *ProductsRepositoryImpl) FindByID(productId int) (product model.Product, err error) {
	var productFound model.Product
	result := p.Db.Find(&productFound, productId)
	if result != nil {
		return productFound, nil

	} else {
		return productFound, errors.New("product not found")
	}
}

// SaveProduct implements repository.ProductsRepository.
func (p *ProductsRepositoryImpl) SaveProduct(product model.Product) {
	result := p.Db.Create(&product)
	helper.ErrorPanic(result.Error)
}

// Update implements repository.ProductsRepository.
func (p *ProductsRepositoryImpl) Update(product model.Product) {
	var updatedProduct = request.UpdateProductsRequest{
		Id:            product.Id,
		Name:          product.Name,
		Category:      product.Category,
		OriginalPrice: product.OriginalPrice,
		DiscountPrice: product.DiscountPrice,
		Supermarket:   product.Supermarket,
	}
	result := p.Db.Model(&product).Updates(updatedProduct)
	helper.ErrorPanic(result.Error)
}

// ClearProducts implements repository.ProductsRepository.
func (p *ProductsRepositoryImpl) ClearProducts() {
	result := p.Db.Exec("DELETE FROM products")
	helper.ErrorPanic(result.Error)
}
