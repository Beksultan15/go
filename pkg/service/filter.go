package service 

import (
	"github.com/Beksultan15/project_go/models"
	"github.com/Beksultan15/project_go/pkg/repository"
	"github.com/gin-gonic/gin"
)

type FilterService struct {
	repo repository.Filtering
}

func NewFilterService(repo repository.Filtering) *FilterService {
	return &FilterService{repo: repo}
}

func (a *FilterService) FilteringProduct(c *gin.Context,lte,gte int) ([]models.Products,error) {
	return a.repo.FilteringProduct(c,lte,gte)
	
}
 