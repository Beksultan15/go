package handler

import (
	"net/http"

	"github.com/Beksultan15/project_go/models"
	"github.com/gin-gonic/gin"
)


type CommentStr struct {
	Product_id int `json:"product_id" binding:"required"`
	User_id int `json:"user_id" binding:"required"`
	Rating int `json:"rating" binding:"required"`
	Comment string `json:"comment" binding:"required"` 
}


func (h *Handler) comment(c *gin.Context){
	var input models.Comment

	if err := c.BindJSON(&input); err!= nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	msg,err := h.services.Commenting.CommentProduct(input)
	if err!= nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"id":msg,
	})
}
