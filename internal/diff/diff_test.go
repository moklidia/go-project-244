package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDiff(t *testing.T) {
	data1 := map[string]interface{}{
		"host":    "hexlet.io",
		"timeout": 50,
		"proxy":   "123.234.53.22",
		"follow":  false,
	}

	data2 := map[string]interface{}{
		"timeout": 20,
		"verbose": true,
		"host":    "hexlet.io",
	}

	var result []Diff
	got := GenerateDiff(data1, data2, &result)

	want := []Diff{
		{Type: Unchanged, Key: "host", Value: "hexlet.io"},
		{Type: Changed, Key: "timeout", OldValue: 50, NewValue: 20},
		{Type: Removed, Key: "proxy", Value: "123.234.53.22"},
		{Type: Removed, Key: "follow", Value: false},
		{Type: Added, Key: "verbose", Value: true},
	}

	assert.ElementsMatch(t, want, got)

}

func TestGenerateNestedDiff(t *testing.T) {
	data1 := map[string]interface{}{
		"host": "hexlet.io",
		"common": map[string]interface{}{
			"setting1": "Value 1",
			"setting2": 200,
		},
	}
	
	data2 := map[string]interface{}{
		"host": "hexlet.io",
		"common": map[string]interface{}{
			"setting1": "Value 1",
			"setting3": true,
		},
	}

	var result []Diff
	got := GenerateDiff(data1, data2, &result)

	want := []Diff{
		{Type: Unchanged, Key: "host", Value: "hexlet.io"},
		{
			Type:     Parent,
			Key:      "common",
			Value:    nil,
			Children: []Diff{
				{Type: Unchanged, Key: "setting1", Value: "Value 1"},
				{Type: Removed, Key: "setting2", Value: 200},
				{Type: Added, Key: "setting3", Value: true},
			},
		},
	}

	assert.ElementsMatch(t, want, got)

}
