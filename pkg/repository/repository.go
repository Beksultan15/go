package repository

import (
	"github.com/Beksultan15/project_go/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username,password string) (models.User, error)
}

type Searching interface {
	GetSearchingProduct(c *gin.Context) ([]models.Products, error)
}

type Filtering interface {
	FilteringProduct(c *gin.Context,lte,gte int) ([]models.Products, error)
}

type CreatingProduct interface {
	CreateProduct(product models.Products) (int,error)
}

type Commenting interface {
	CommentProduct(comment models.Comment) ( int, error)
}
type Purchasing interface{
	PurchasingItems(id,userId int) (string,error)
}

type Repository struct {
	Authorization
	Searching
	Filtering
	Commenting
	CreatingProduct
	Purchasing
}

func NewRepository(db *sqlx.DB) *Repository {
    return &Repository{
		Authorization: NewAuthPostgres(db),
		Searching: NewSearchPostgres(db),
		Filtering: NewFilterPostgres(db),
		Commenting: NewCommentPostgres(db),
		CreatingProduct: NewProductPostgres(db),
		Purchasing: NewPurchasingProduct(db),
	}
}