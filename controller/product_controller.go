package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//Usecase
}

// Função para inicializar esta estrutura
func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			Id:    1,
			Name:  "Sashimi",
			Price: 75,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
