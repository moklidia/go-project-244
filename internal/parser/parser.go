package parser

import "encoding/json"

func ParseJson(data string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}

	return result, nil
}
