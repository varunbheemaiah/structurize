package view

import (
	"converter/internal/pkg/json2schema"
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
