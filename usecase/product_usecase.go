package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductById(productId int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(productId)
	if err != nil {
		return nil, nil
	}

	return product, nil
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.Id = productId

	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(productId int, product model.Product) (model.Product, error) {

	product, err := pu.repository.UpdateProduct(productId, product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProduct(productId int) (int, error) {
	rowsAffected, err := pu.repository.DeleteProduct(productId)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
