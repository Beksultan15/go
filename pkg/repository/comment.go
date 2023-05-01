package repository

import (
	"fmt"
	"github.com/Beksultan15/project_go/models"
	"github.com/jmoiron/sqlx"
)


type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres{
	return &CommentPostgres{db:db}
}


func (a *CommentPostgres) CommentProduct(comment models.Comment) (int,error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s(id, user_id, product_id, rating,comment_body) VALUES ($1,$2,$3,$4,$5) RETURNING id", comment_table)
	row := a.db.QueryRow(query,comment.Id,comment.User_id,comment.Product_id,comment.Rating,comment.Comment)
	
	if err:= row.Scan(&id); err!= nil {
		return 0,err
	}

	return id,nil
}