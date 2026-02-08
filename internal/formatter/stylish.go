package formatter

import (
	"code/internal/diff"
	"fmt"
	"sort"
	"strings"
)

func formatStylish(data []diff.Diff, depth int) []string {
	var lines []string
	baseIndent := strings.Repeat(" ", indentSize*depth)

	for _, item := range data {
		switch item.Type {
		case diff.Parent:
			lines = append(lines, fmt.Sprintf("    %s%s: {", baseIndent, item.Key))
			lines = append(lines, formatStylish(item.Children, depth+1)...)
			lines = append(lines, fmt.Sprintf("    %s}", baseIndent))
		case diff.Added:
			if len(item.Children) > 0 {
				lines = append(lines, fmt.Sprintf("  %s+ %s: {", baseIndent, item.Key))
				lines = append(lines, formatStylishValueOnly(item.Children, depth+1)...)
				lines = append(lines, fmt.Sprintf("    %s}", baseIndent))
			} else {
				lines = append(lines, formatStylishLine("+", item.Key, item.Value, baseIndent))
			}
		case diff.Removed:
			if len(item.Children) > 0 {
				lines = append(lines, fmt.Sprintf("  %s- %s: {", baseIndent, item.Key))
				lines = append(lines, formatStylishValueOnly(item.Children, depth+1)...)
				lines = append(lines, fmt.Sprintf("    %s}", baseIndent))
			} else {
				lines = append(lines, formatStylishLine("-", item.Key, item.Value, baseIndent))
			}
		case diff.Unchanged:
			newLine := formatStylishLine(" ", item.Key, item.Value, baseIndent)
			lines = append(lines, newLine)
	case diff.Changed:
		oldValue := formatValueToStylish(item.OldValue, depth)
		newValue := formatValueToStylish(item.NewValue, depth)
		lines = append(lines, fmt.Sprintf("  %s- %s: %s", baseIndent, item.Key, oldValue))
		lines = append(lines, fmt.Sprintf("  %s+ %s: %s", baseIndent, item.Key, newValue))
		}
	}

	return lines
}

func formatStylishLine(prefix, key string, value interface{}, indent string) string {
	return fmt.Sprintf("  %s%s %s: %s", indent, prefix, key, formatValue(value))
}

func formatStylishValueOnly(data []diff.Diff, depth int) []string {
	var lines []string
	keyIndent := strings.Repeat(" ", indentSize*(depth+1))

	for _, item := range data {
		if len(item.Children) > 0 {
			lines = append(lines, fmt.Sprintf("%s%s: {", keyIndent, item.Key))
			lines = append(lines, formatStylishValueOnly(item.Children, depth+1)...)
			lines = append(lines, fmt.Sprintf("%s}", keyIndent))
		} else {
			lines = append(lines, fmt.Sprintf("%s%s: %s", keyIndent, item.Key, formatValue(item.Value)))
		}
	}
	return lines
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

func formatValueToStylish(value interface{}, depth int) string {
	if value == nil {
		return "null"
	}
	if m, ok := value.(map[string]interface{}); ok {
		return mapToStylish(m, depth)
	}
	return fmt.Sprintf("%v", value)
}

func mapToStylish(m map[string]interface{}, depth int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	keyIndent := strings.Repeat(" ", indentSize*(depth+2))
	closeIndent := strings.Repeat(" ", indentSize*(depth+1))
	
	var lines []string
	for _, k := range keys {
		v := m[k]
		lines = append(lines, fmt.Sprintf("%s%s: %s", keyIndent, k, formatValueToStylish(v, depth+1)))
	}
	result := strings.Join(lines, "\n")
	return fmt.Sprintf("{\n%s\n%s}", result, closeIndent)
}
