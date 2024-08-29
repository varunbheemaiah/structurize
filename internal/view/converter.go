package view

import (
	"converter/internal/pkg/json2schema"
	"converter/internal/pkg/json2xml"
	"converter/internal/pkg/xml2json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p Provider) ConvertJSONToSchema(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonStr := string(body)

	schema := json2schema.ConvertJSONToSchema(jsonStr)

	c.JSON(200, schema)

}

func (p Provider) ConvertJSONToXML(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonStr := string(body)

	xmlStr := json2xml.ConvertJSONToXML(jsonStr)

	c.String(200, xmlStr)

}

func (p Provider) ConvertXMLToJSON(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	xmlStr := string(body)

	jsonStr, err := xml2json.ConvertXMLToJSON(xmlStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, jsonStr)

}
