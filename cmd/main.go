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
	router.POST("/json-2-xml", converterProvider.ConvertJSONToXML)
	router.POST("/xml-2-json", converterProvider.ConvertXMLToJSON)
	router.POST("/json-2-struct", converterProvider.ConvertJSONToStruct)

	router.Run(":8000")

}
