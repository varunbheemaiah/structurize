package json2struct

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ConvertJSONToSchema(jsonStr string, bson, xml, defaultVal, omitempty bool) (string, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	var sb strings.Builder
	sb.WriteString("type Generated struct {\n")

	for k, v := range jsonData {
		fieldName := capitalizeFirstLetter(k)
		fieldType := inferGoType(v)

		jsonTag := fmt.Sprintf(`json:"%s`, k)
		if omitempty {
			jsonTag += ",omitempty"
		}
		jsonTag += `"`

		bsonTag := ""
		if bson {
			bsonTag = fmt.Sprintf(`bson:"%s`, k)
			if omitempty {
				bsonTag += ",omitempty"
			}
			bsonTag += `"`
		}

		xmlTag := ""
		if xml {
			xmlTag = fmt.Sprintf(`xml:"%s`, k)
			if omitempty {
				xmlTag += ",omitempty"
			}
			xmlTag += `"`
		}

		defaultTag := ""
		if defaultVal {
			defaultStr := fmt.Sprintf("%v", v)
			defaultTag = fmt.Sprintf(`default="%s"`, defaultStr)
		}

		allTags := combineTags(jsonTag, bsonTag, xmlTag, defaultTag)

		sb.WriteString(fmt.Sprintf("    %s %s %s\n", fieldName, fieldType, allTags))
	}

	sb.WriteString("}\n")

	return sb.String(), nil
}

func capitalizeFirstLetter(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

func inferGoType(val interface{}) string {
	switch v := val.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		if float64(int64(v)) == v {
			return "int"
		}
		return "float64"
	default:
		return "string"
	}
}

func combineTags(tags ...string) string {
	nonEmptyTags := []string{}
	for _, t := range tags {
		if t != "" {
			nonEmptyTags = append(nonEmptyTags, t)
		}
	}

	if len(nonEmptyTags) == 0 {
		return ""
	}

	return fmt.Sprintf("`%s`", strings.Join(nonEmptyTags, " "))
}
