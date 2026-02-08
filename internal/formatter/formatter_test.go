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

func TestFormatNestedAdded(t *testing.T) {
	data := []diff.Diff{
		{
			Type: diff.Added,
			Key:  "group3",
			Children: []diff.Diff{
				{Type: diff.Added, Key: "fee", Value: 100500},
				{Type: diff.Added, Key: "nested", Value: "value"},
			},
		},
	}

	want := `{
  + group3: {
        fee: 100500
        nested: value
    }
}`

	got := Format(data, "stylish")
	assert.Equal(t, want, got)
}

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
