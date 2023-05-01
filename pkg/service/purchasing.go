package service

import "github.com/Beksultan15/project_go/pkg/repository"

type PurchaseService struct {
	repo repository.Purchasing
}

func NewPurchasingService(repo repository.Purchasing) * PurchaseService{
	return &PurchaseService{repo: repo}
}

func (a *PurchaseService) PurchasingItems(id,userId int) (string,error){
	return a.repo.PurchasingItems(id,userId)
}