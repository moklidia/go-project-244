package formatter

import (
	"code/internal/diff"
	"fmt"

)

func formatPlain(data []diff.Diff, prefix string) []string {
	var lines []string
	for _, item := range data {
		fullKey := item.Key
		if prefix != "" {
			fullKey = prefix + "." + item.Key
		}
		switch item.Type {
		case diff.Parent:
			childLines := formatPlain(item.Children, fullKey)
			lines = append(lines, childLines...)
		case diff.Added:
			if len(item.Children) > 0 {
				lines = append(lines, fmt.Sprintf("Property '%s' was added with value: [complex value]", fullKey))
			} else {
				lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", fullKey, formatPlainValue(item.Value)))
			}
		case diff.Removed:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", fullKey))
		case diff.Changed:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", fullKey, formatPlainValue(item.OldValue), formatPlainValue(item.NewValue)))
		}
	}

	return lines
}

func formatPlainValue(v interface{}) string {
	if v == nil {
		return "null"
	}
	if _, ok := v.(map[string]interface{}); ok {
		return "[complex value]"
	}
	if s, ok := v.(string); ok {
		return fmt.Sprintf("'%s'", s)
	}
	return fmt.Sprintf("%v", v)
}
