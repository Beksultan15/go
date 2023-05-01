package handler

import (
	// "net/http"

	"github.com/Beksultan15/project_go/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
    return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	 {
		auth.POST("/signup",h.signup)
		auth.POST("/signin",h.signin)
	 }

	 product := router.Group("/product")
	 {
		product.GET("/search/",h.Search)
		product.GET("/filter/",h.filter)
		product.POST("/comment/",h.comment)
		product.POST("/create",h.createproduct)
		product.POST("/purchasing",h.purchase)
	 }

	 return router
}