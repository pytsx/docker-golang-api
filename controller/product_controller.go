package controller

import (
	"go-api/model"
	usecase "go-api/useCase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type produceController struct {
	// usecase
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) produceController {
	return produceController{
		productUsecase: usecase,
	}
}

func (p *produceController) GetProducts(ctx *gin.Context) {
	// usecase

	proucts, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, proucts)
}

func (p *produceController) CreateProduct(ctx *gin.Context) {
	// usecase
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return 
	}
	
	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
} 

func (p *produceController) GetProductById(ctx *gin.Context) {
	// usecase
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Product ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Product ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return 
	}

	ctx.JSON(http.StatusOK, product)
}