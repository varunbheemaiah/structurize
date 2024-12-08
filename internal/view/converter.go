package view

import (
	"converter/internal/pkg/json2schema"
	"converter/internal/pkg/json2struct"
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

func (p Provider) ConvertJSONToStruct(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the values of xml, json, bson, omitempty and default from the request parameters
	xml := c.Query("xml")
	bson := c.Query("bson")
	omitempty := c.Query("omitempty")
	defaultValue := c.Query("default")

	// Convert the xml, json, bson, omitempty and default values to boolean
	xmlBool := xml == "true"
	bsonBool := bson == "true"
	omitemptyBool := omitempty == "true"
	defaultBool := defaultValue == "true"

	jsonStr := string(body)

	goStruct, err := json2struct.ConvertJSONToSchema(jsonStr, bsonBool, xmlBool, defaultBool, omitemptyBool)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(goStruct))

}
