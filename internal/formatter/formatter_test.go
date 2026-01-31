package formatter

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	data := []diff.Diff{
		{Type: diff.Removed, Key: "follow", Value: false},
		{Type: diff.Unchanged, Key: "host", Value: "hexlet.io"},
		{Type: diff.Removed, Key: "proxy", Value: "123.234.53.22"},
		{Type: diff.Changed, Key: "timeout", OldValue: 50, NewValue: 20},
		{Type: diff.Added, Key: "verbose", Value: true},
	}

	want := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	got := Format(data, "stylish")

	assert.Equal(t, want, got)
}

func TestFormatNested(t *testing.T) {
	// Данные как возвращает GenerateDiff (дерево с Parent и Children)
	data := []diff.Diff{
		{Type: diff.Unchanged, Key: "host", Value: "hexlet.io"},
		{
			Type:  diff.Parent,
			Key:   "common",
			Children: []diff.Diff{
				{Type: diff.Unchanged, Key: "setting1", Value: "Value 1"},
				{Type: diff.Removed, Key: "setting2", Value: 200},
				{Type: diff.Added, Key: "setting3", Value: true},
			},
		},
	}

	want := `{
    host: hexlet.io
  common: {
      setting1: Value 1
    - setting2: 200
    + setting3: true
  }
}`

	got := Format(data, "stylish")

	assert.Equal(t, want, got)
}
