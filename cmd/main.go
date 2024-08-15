package main

import (
	"converter/internal/view"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	converterProvider := view.NewProvider()
	router.POST("/json-2-schema", converterProvider.ConvertJSONToSchema)

	router.Run(":8000")

}
