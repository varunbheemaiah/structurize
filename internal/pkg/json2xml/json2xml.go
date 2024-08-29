package json2xml

import (
	"encoding/json"
	"fmt"
	"log"
)

func ConvertJSONToXML(jsonString string) string {
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonData)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	var xmlString string
	switch data := jsonData.(type) {
	case map[string]interface{}:
		xmlString = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><root>" + mapToXML(data) + "</root>"
	case []interface{}:
		xmlString = "<?xml version=\"1.0\" encoding=\"UTF-8\"?><root>" + arrayToXML(data) + "</root>"
	default:
		log.Fatalf("Unsupported JSON format")
	}

	return xmlString
}

func mapToXML(data map[string]interface{}) string {
	xmlStr := ""

	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			xmlStr += fmt.Sprintf("<%s>%s</%s>", key, mapToXML(v), key)
		case []interface{}:
			xmlStr += fmt.Sprintf("<%s>%s</%s>", key, arrayToXML(v), key)
		default:
			xmlStr += fmt.Sprintf("<%s>%v</%s>", key, v, key)
		}
	}

	return xmlStr
}

func arrayToXML(data []interface{}) string {
	xmlStr := ""

	for _, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			xmlStr += "<item>" + mapToXML(v) + "</item>"
		case []interface{}:
			xmlStr += "<item>" + arrayToXML(v) + "</item>"
		default:
			xmlStr += fmt.Sprintf("<item>%v</item>", v)
		}
	}

	return xmlStr
}
