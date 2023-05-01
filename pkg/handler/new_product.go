package handler

import (
	"net/http"

	"github.com/Beksultan15/project_go/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createproduct(c *gin.Context){

	var input models.Products

	if err := c.BindJSON(&input); err!= nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	msg,err := h.services.CreatingProduct.CreateProduct(input)
	if err!= nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"id":msg,
	})
}