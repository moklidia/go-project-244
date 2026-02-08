package diff

import (
	"sort"
)

type ChangeType string

const (
	Added     ChangeType = "added"
	Removed   ChangeType = "removed"
	Changed   ChangeType = "changed"
	Unchanged ChangeType = "unchanged"
	Parent    ChangeType = "parent" // узел с детьми, как в PHP
)

type Diff struct {
	Type     ChangeType
	Key      string
	Value    interface{}
	OldValue interface{}
	NewValue interface{}
	Children []Diff
}

func GenerateDiff(data1, data2 map[string]interface{}, diff *[]Diff) []Diff {
	allKeys := make(map[string]bool)

	for key := range data1 {
		allKeys[key] = true
	}

	for key := range data2 {
		allKeys[key] = true
	}

	for key := range allKeys {
		value1, ok1 := data1[key]
		value2, ok2 := data2[key]
		if !ok1 {
			if isNested(value2) {
				map2 := value2.(map[string]interface{})
				var children []Diff
				GenerateDiff(map[string]interface{}{}, map2, &children)
				*diff = append(*diff, Diff{Type: Added, Key: key, Children: children })
			} else {
				addAdded(diff, key, value2)
			}
			continue
		}
		if !ok2 {
			if isNested(value1) {
				map1 := value1.(map[string]interface{})
				var children []Diff
				GenerateDiff(map1, map[string]interface{}{}, &children)
				*diff = append(*diff, Diff{Type: Removed, Key: key, Children: children})
			} else {
				addRemoved(diff, key, value1)
			}
			continue
		}
		if isNested(value1) && isNested(value2) {
			children := generateNestedDiff(value1, value2)
			*diff = append(*diff, Diff{Type: Parent, Key: key, Children: children})
			continue
		}
		if value1 == value2 {
			addUnchanged(diff, key, value1)
			continue
		}
		addChanged(diff, key, value1, value2)
	}

	sort.Slice(*diff, func(a, b int) bool {
		return (*diff)[a].Key < (*diff)[b].Key
	})

	return *diff
}

func addUnchanged(diff *[]Diff, key string, value interface{}) {
	newEntry := Diff{Type: Unchanged, Key: key, Value: value}
	*diff = append(*diff, newEntry)
}

func addChanged(diff *[]Diff, key string, value1, value2 interface{}) {
	newEntry := Diff{Type: Changed, Key: key, OldValue: value1, NewValue: value2}
	*diff = append(*diff, newEntry)
}

func addRemoved(diff *[]Diff, key string, value interface{}) {
	newEntry := Diff{Type: Removed, Key: key, Value: value}
	*diff = append(*diff, newEntry)
}

func addAdded(diff *[]Diff, key string, value interface{}) {
	newEntry := Diff{Type: Added, Key: key, Value: value}
	*diff = append(*diff, newEntry)
}

func isNested(value interface{}) bool {
	_, ok := value.(map[string]interface{})

	return ok
}

func generateNestedDiff(value1, value2 interface{}) []Diff {
	map1, _ := value1.(map[string]interface{})
	map2, _ := value2.(map[string]interface{})
	var children []Diff
	GenerateDiff(map1, map2, &children)
	return children
}
