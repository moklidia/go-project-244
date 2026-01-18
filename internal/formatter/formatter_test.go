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
