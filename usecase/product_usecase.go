package usecase

import (
	"go-api/model"
)
type ProductUsecase struct {
}

func NewProductUsecase() *ProductUsecase {
	return &ProductUsecase{}
}

func (p *ProductUsecase) GetAllProducts() ([]model.Product, error) {
	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}
	return []model.Product(products), nil
}