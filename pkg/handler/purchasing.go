package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Purch struct {
	id string `json:"id" binding:"required"`
	userId string `json:"user_id" binding:"required"`
}

func (h *Handler) purchase(c *gin.Context){
	
	var input Purch
	// var test2 test
	if err := c.BindJSON(&input); err!= nil {
			newErrorResponse(c,http.StatusBadRequest, err.Error())
			return
	}
	fmt.Println(input.id)

	msg,err := h.services.Purchasing.PurchasingItems(2,2)

	if err!= nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"msg":msg,
		"status":http.StatusOK,
	})

}