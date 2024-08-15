package view

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProviderInterface interface {
	ConvertJSONToSchema(c *gin.Context)
}

type Provider struct {
	v *validator.Validate
}

func NewProvider() ProviderInterface {
	return &Provider{
		v: validator.New(),
	}
}