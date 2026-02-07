package main

import (
	"path/filepath"
	"testing"
	"code"

	"github.com/stretchr/testify/assert"
)

func TestRunJson(t *testing.T) {
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

	got, err := code.GenDiff(file1, file2, "stylish")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

func TestRunYml(t *testing.T) {
	file1 := filepath.Join("..", "..", "testdata/fixtures/file1.yml")
	file2 := filepath.Join("..", "..", "testdata/fixtures/file2.yml")

	want := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	got, err := code.GenDiff(file1, file2, "stylish")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}
