package service

import (
	"github.com/Beksultan15/project_go/models"
	"github.com/Beksultan15/project_go/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	RefreshToken(username,password string) (string,error)
}

type Searhing interface {
	GetSearchingProduct(c *gin.Context) ([]models.Products,error)
}


type Filtering interface{
	FilteringProduct(c *gin.Context,lte,gte int) ([]models.Products,error)
}

type CreatingProduct interface{
	CreateProduct(product models.Products) (int,error)
}

type Commenting interface{
	CommentProduct(models.Comment) (int,error)
}

type Purchasing interface {
	PurchasingItems(id,userId int) (string,error)
}

type Service struct {
	Authorization
	Searhing
	Filtering
	Commenting
	CreatingProduct
	Purchasing
}

func NewService(repos *repository.Repository) *Service {
    return &Service{
		Authorization: newAuthService(repos.Authorization),
		Searhing: newSearchService(repos.Searching),
		Filtering: NewFilterService(repos.Filtering),
		Commenting: NewCommentService(repos.Commenting),
		CreatingProduct: NewProductService(repos.CreatingProduct),
		Purchasing: NewPurchasingService(repos.Purchasing),
	}  
}