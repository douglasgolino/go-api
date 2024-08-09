package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

// Função para inicializar esta estrutura
func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()
	if err != nil {
		response := model.Response{
			Message: "Erro ao processar a solicitação",
		}

		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("id_product")
	if id == "" {
		response := model.Response{
			Message: "Informe o Id do produto",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número inteiro",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		response := model.Response{
			Message: "Erro ao processar a solicitação",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		response := model.Response{
			Message: "Solicitação Inválida " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		response := model.Response{
			Message: "Erro ao processar a solicitação",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {

	id := ctx.Param("id_product")
	if id == "" {
		response := model.Response{
			Message: "Informe o Id do produto",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número inteiro",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product

	err = ctx.BindJSON(&product)
	if err != nil {
		response := model.Response{
			Message: "Solicitação Inválida " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updatedProduct, err := p.productUsecase.UpdateProduct(productId, product)

	if err != nil {
		response := model.Response{
			Message: "Erro ao processar a solicitação",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {

	id := ctx.Param("id_product")
	if id == "" {
		response := model.Response{
			Message: "Informe o Id do produto",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número inteiro",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	rowsAffected, err := p.productUsecase.DeleteProduct(productId)
	if err != nil {
		response := model.Response{
			Message: "Erro ao processar a solicitação",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if rowsAffected == 0 {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := model.Response{
		Message: "Produto deletado com sucesso",
	}

	ctx.JSON(http.StatusOK, response)
}
