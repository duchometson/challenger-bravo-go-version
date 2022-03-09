package currency

import "github.com/gin-gonic/gin"

type Currency struct {
	service Service
}

func Get(ctx *gin.Context) {}

func Add(ctx *gin.Context) {}

func Delete(ctx *gin.Context) {}

func Convert(ctx *gin.Context) {}

func New(service Service) *Currency {
	return &Currency{
		service: service,
	}
}
