package service

import (
	"github.com/Beksultan15/project_go/models"
	"github.com/Beksultan15/project_go/pkg/repository"
)

type ProductService struct{
	repo repository.CreatingProduct
}

func NewProductService(repo repository.CreatingProduct) *ProductService{
	return &ProductService{repo:repo}
}

func (a *ProductService) CreateProduct(product models.Products) (int,error){
	return a.repo.CreateProduct(product)
}