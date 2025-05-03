package usecase

import (
	"go-api/model"
	"go-api/repository"
)
type ProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		productRepository: repository,
	}
}

func (p *ProductUsecase) GetAllProducts() ([]model.Product, error) {
	products, err := p.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}