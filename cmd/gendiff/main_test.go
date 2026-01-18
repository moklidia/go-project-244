package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file1 := filepath.Join("..", "..", "testdata/fixtures/file1.json")
	file2 := filepath.Join("..", "..", "testdata/fixtures/file2.json")

	want := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	got, err := run(file1, file2, "stylish")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}
