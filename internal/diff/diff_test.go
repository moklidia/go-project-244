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

	got := GenerateDiff(data1, data2)

	want := []Diff{
		{Type: Unchanged, Key: "host", Value: "hexlet.io"},
		{Type: Changed, Key: "timeout", OldValue: 50, NewValue: 20},
		{Type: Removed, Key: "proxy", Value: "123.234.53.22"},
		{Type: Removed, Key: "follow", Value: false},
		{Type: Added, Key: "verbose", Value: true},
	}

	assert.ElementsMatch(t, want, got)

}
