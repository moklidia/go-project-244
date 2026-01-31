package formatter

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

func Format(data []diff.Diff, _format string) string {
	lines := convert(data, 1)

	result := strings.Join(lines, "\n")
	return fmt.Sprintf("{\n%s\n}", result)
}

func convert(data []diff.Diff, depth int) []string {
	var lines []string

	for _, item := range data {
		switch item.Type {
		case diff.Parent:
			indent := generateIndent(depth)
			lines = append(lines, fmt.Sprintf("%s%s: {", indent, item.Key))
			lines = append(lines, convert(item.Children, depth+1)...)
			lines = append(lines, fmt.Sprintf("%s}", indent))
		case diff.Added:
			if len(item.Children) > 0 {
				indent := generateIndent(depth)
				lines = append(lines, fmt.Sprintf("%s+ %s: {", indent, item.Key))
				lines = append(lines, convert(item.Children, depth+1)...)
				lines = append(lines, fmt.Sprintf("%s}", indent))
			} else {
				lines = append(lines, formatLine("+", item.Key, item.Value, depth))
			}
		case diff.Removed:
			if len(item.Children) > 0 {
				indent := generateIndent(depth)
				lines = append(lines, fmt.Sprintf("%s- %s: {", indent, item.Key))
				lines = append(lines, convert(item.Children, depth+1)...)
				lines = append(lines, fmt.Sprintf("%s}", indent))
			} else {
				lines = append(lines, formatLine("-", item.Key, item.Value, depth))
			}
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
	return fmt.Sprintf("%s%s %s: %s", indent, prefix, key, formatValue(value))
}

func formatValue(v interface{}) string {
	if v == nil {
		return "null"
	}
	if _, ok := v.(map[string]interface{}); ok {
		return "[complex value]"
	}
	return fmt.Sprintf("%v", v)
}

func generateIndent(depth int) string {
	spaces := depth * indentSize
	return strings.Repeat(" ", spaces)
}
