package formatter

import (
	"encoding/json"
	"code/internal/diff"
)

func formatJson(data []diff.Diff) string {
	result, _ := json.Marshal(data)
	return string(result)
}
