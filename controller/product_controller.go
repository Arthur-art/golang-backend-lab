package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) GetAllProducts(ctx *gin.Context) {

	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}

	ctx.JSON(http.StatusOK, products)
}