package service 

import (
	"github.com/Beksultan15/project_go/models"
	"github.com/Beksultan15/project_go/pkg/repository"
	
)

type CommentService struct {
	repo repository.Commenting
}

func NewCommentService(repo repository.Commenting) *CommentService {
	return &CommentService{repo: repo}
}

func (a *CommentService) CommentProduct(comment models.Comment) (int,error) {
	return a.repo.CommentProduct(comment)
}
 