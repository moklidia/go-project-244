package formatter

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

const (
	indentSize = 4
)

func Format(data []diff.Diff, format string) string {
	var result string
	switch format {
	case "stylish":
	  lines := formatStylish(data, 0)
		result = fmt.Sprintf("{\n%s\n}", strings.Join(lines, "\n"))
	case "plain":
		lines := formatPlain(data, "")
		result = strings.Join(lines, "\n")
	case "json":
		result = formatJson(data)
	}

	return result
}
