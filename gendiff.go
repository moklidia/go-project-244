package code

import (
	"code/internal/diff"
	"code/internal/formatter"
	"code/internal/parser"
	"fmt"
	"os"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := os.ReadFile(filepath1)
	if err != nil {
		return "", err
	}
	data2, err := os.ReadFile(filepath2)
	if err != nil {
		return "", err
	}
	parsed1, err := parser.Parse(string(data1))
	if err != nil {
		return "", fmt.Errorf("error parsing %s: %w", filepath1, err)
	}
	parsed2, err := parser.Parse(string(data2))
	if err != nil {
		return "", fmt.Errorf("error parsing %s: %w", filepath2, err)
	}
	var diffData []diff.Diff
	diffData = diff.GenerateDiff(parsed1, parsed2, &diffData)
	return formatter.Format(diffData, format), nil
}
