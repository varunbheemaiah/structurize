package xml2json

import (
	"encoding/xml"
	"strings"
)

type Node struct {
	XMLName xml.Name
	Content string     `xml:",chardata"`
	Nodes   []Node     `xml:",any"`
	Attrs   []xml.Attr `xml:",attr"`
}

func ConvertXMLToJSON(xmlString string) (map[string]interface{}, error) {
	var node Node
	err := xml.Unmarshal([]byte(xmlString), &node)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	nodeToMap(node, result)

	return result, nil
}

func nodeToMap(node Node, m map[string]interface{}) {
	if len(node.Attrs) > 0 {
		attrMap := make(map[string]string)
		for _, attr := range node.Attrs {
			attrMap[attr.Name.Local] = attr.Value
		}
		m["#attrs"] = attrMap
	}

	for _, child := range node.Nodes {
		childMap := make(map[string]interface{})
		nodeToMap(child, childMap)

		if existing, found := m[child.XMLName.Local]; found {
			switch existing := existing.(type) {
			case []interface{}:
				m[child.XMLName.Local] = append(existing, childMap)
			default:
				m[child.XMLName.Local] = []interface{}{existing, childMap}
			}
		} else {
			m[child.XMLName.Local] = childMap
		}
	}

	if len(strings.TrimSpace(node.Content)) > 0 {
		m["#text"] = node.Content
	}
}
