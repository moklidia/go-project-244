package formatter

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

func Format(data []diff.Diff) string {
	lines := convert(data, 1)

	result := strings.Join(lines, "\n")
	return fmt.Sprintf("{\n%s\n}", result)
}

func convert(data []diff.Diff, depth int) []string {
	var lines []string

	for _, item := range data {
		switch item.Type {
		case diff.Added:
			newLine := formatLine("+", item.Key, item.Value, depth)
			lines = append(lines, newLine)
		case diff.Removed:
			newLine := formatLine("-", item.Key, item.Value, depth)
			lines = append(lines, newLine)
		case diff.Unchanged:
			newLine := formatLine(" ", item.Key, item.Value, depth)
			lines = append(lines, newLine)
		case diff.Changed:
			oldLine := formatLine("-", item.Key, item.OldValue, depth)
			newLine := formatLine("+", item.Key, item.NewValue, depth)
			lines = append(lines, oldLine)
			lines = append(lines, newLine)
		}
	}

	return lines
}

func formatLine(prefix, key string, value interface{}, depth int) string {
	indent := generateIndent(depth)
	return fmt.Sprintf("%s%s %s: %v", indent, prefix, key, value)
}

func generateIndent(depth int) string {
	spaces := depth * indentSize
	return strings.Repeat(" ", spaces)
}
