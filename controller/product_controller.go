package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//Usecase
}

func NewProductController() *productController {
	return &productController{
		//Usecase: usecase,
	}
}

func (p *productController) GetAllProducts(ctx *gin.Context) {

	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}

	ctx.JSON(http.StatusOK, products)
}