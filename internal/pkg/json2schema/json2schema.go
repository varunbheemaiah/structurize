package json2schema

import (
	"encoding/json"
	"log"
	"reflect"
)

func ConvertJSONToSchema(jsonString string) map[string]interface{} {

	var jsonData interface{}
	if err := json.Unmarshal([]byte(jsonString), &jsonData); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	schema := generateJSONSchema(jsonData)

	return schema

}

func generateJSONSchema(data interface{}) map[string]interface{} {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		properties := make(map[string]interface{})
		for key, value := range data.(map[string]interface{}) {
			properties[key] = generateJSONSchema(value)
		}
		return map[string]interface{}{
			"type":       "object",
			"properties": properties,
		}
	case reflect.Slice:
		arrayItems := data.([]interface{})
		var itemSchema interface{}
		if len(arrayItems) > 0 {
			itemSchema = generateJSONSchema(arrayItems[0])
		} else {
			itemSchema = map[string]interface{}{
				"type": "string",
			}
		}
		return map[string]interface{}{
			"type":  "array",
			"items": itemSchema,
		}
	case reflect.String:
		return map[string]interface{}{
			"type": "string",
		}
	case reflect.Float64:
		return map[string]interface{}{
			"type": "number",
		}
	case reflect.Bool:
		return map[string]interface{}{
			"type": "boolean",
		}
	case reflect.Int:
		return map[string]interface{}{
			"type": "integer",
		}
	default:
		return map[string]interface{}{
			"type": "string",
		}
	}
}
