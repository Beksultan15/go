package service

import (
	"github.com/Beksultan15/project_go/models"
	"github.com/Beksultan15/project_go/pkg/repository"
	"github.com/gin-gonic/gin"
)

type SearchService struct {
	repo repository.Searching
}

func newSearchService(repo repository.Searching) *SearchService{
	return &SearchService{repo:repo}
}

func (a *SearchService) GetSearchingProduct(c *gin.Context) ([]models.Products,error) {
	return a.repo.GetSearchingProduct(c)
	
}
