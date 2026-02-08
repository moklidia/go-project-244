package formatter

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatPlain(t *testing.T) {
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
		{Type: diff.Changed, Key: "port", OldValue: 80, NewValue: 3000},
		{
			Type: diff.Added,
			Key:  "group3",
			Children: []diff.Diff{
				{Type: diff.Added, Key: "fee", Value: 100500},
			},
		},
	}

	want := `Property 'common.setting2' was removed
Property 'common.setting3' was added with value: true
Property 'port' was updated. From 80 to 3000
Property 'group3' was added with value: [complex value]`

	got := Format(data, "plain")

	assert.Equal(t, want, got)
}
