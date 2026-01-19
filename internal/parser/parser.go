package parser

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func Parse(data string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err == nil {
		return result, nil
	}

	if err := yaml.Unmarshal([]byte(data), &result); err == nil {
		return result, nil
	}

	return nil, fmt.Errorf("unable to parse as JSON or YAML")
}
