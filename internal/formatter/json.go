package formatter

import (
	"encoding/json"
	"code/internal/diff"
)

func formatJson(_data []diff.Diff) string {
	return `{"test": "value"}`
}
