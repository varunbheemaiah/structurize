package json2struct

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConvertJSONToSchema takes a map representing a JSON object and returns a Go struct definition.
// The parameters bson, xml, defaultVal, omitempty determine which tags to include in the struct fields.
func ConvertJSONToSchema(jsonStr string, bson, xml, defaultVal, omitempty bool) (string, error) {
	// Parse the JSON into a map
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Start building the struct definition
	var sb strings.Builder
	sb.WriteString("type Generated struct {\n")

	for k, v := range jsonData {
		// Determine the Go field name (simple approach: capitalize the first letter)
		fieldName := capitalizeFirstLetter(k)
		fieldType := inferGoType(v)

		// Build the tag for json
		jsonTag := fmt.Sprintf(`json:"%s`, k)
		if omitempty {
			jsonTag += ",omitempty"
		}
		jsonTag += `"`

		// Add optional bson tag
		bsonTag := ""
		if bson {
			bsonTag = fmt.Sprintf(`bson:"%s`, k)
			if omitempty {
				bsonTag += ",omitempty"
			}
			bsonTag += `"`
		}

		// Add optional xml tag
		xmlTag := ""
		if xml {
			xmlTag = fmt.Sprintf(`xml:"%s`, k)
			if omitempty {
				xmlTag += ",omitempty"
			}
			xmlTag += `"`
		}

		// Add optional default tag
		defaultTag := ""
		if defaultVal {
			// Convert the default value to string form
			defaultStr := fmt.Sprintf("%v", v)
			// If it's a string, we don't need extra quotes in the tag,
			// but typically for tags, it's just a literal value.
			defaultTag = fmt.Sprintf(`default="%s"`, defaultStr)
		}

		// Combine tags
		allTags := combineTags(jsonTag, bsonTag, xmlTag, defaultTag)

		// Add the field line
		sb.WriteString(fmt.Sprintf("    %s %s %s\n", fieldName, fieldType, allTags))
	}

	sb.WriteString("}\n")

	return sb.String(), nil
}

// capitalizeFirstLetter capitalizes just the first letter of a string.
func capitalizeFirstLetter(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

// inferGoType attempts to infer a Go type from an interface{} value from JSON.
func inferGoType(val interface{}) string {
	switch v := val.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		// Check if the float64 is actually an integer value
		if float64(int64(v)) == v {
			return "int"
		}
		return "float64"
	default:
		// If needed, handle nested objects or arrays here.
		// For now, default to string.
		return "string"
	}
}

// combineTags combines multiple tags into a single backquoted string.
func combineTags(tags ...string) string {
	// Filter out empty tags and combine them
	nonEmptyTags := []string{}
	for _, t := range tags {
		if t != "" {
			// Ensure there's a space between tags to separate them.
			nonEmptyTags = append(nonEmptyTags, t)
		}
	}

	if len(nonEmptyTags) == 0 {
		return ""
	}

	// Combine all tags into a single backtick-delimited string
	return fmt.Sprintf("`%s`", strings.Join(nonEmptyTags, " "))
}
