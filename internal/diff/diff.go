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
)

type Diff struct {
	Type     ChangeType
	Key      string
	Value    interface{}
	OldValue interface{}
	NewValue interface{}
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
		if ok1 && ok2 {
			if isBothNested(value1, value2) {
				addNestedDiff(key, value1, value2, diff)
				continue
			}
			if value1 == value2 {
				addUnchanged(diff, key, value1)
			} else {
				addChanged(diff, key, value1, value2)
			}
		}
		if !ok1 {
			addAdded(diff, key, value2)
		}

		if !ok2 {
			addRemoved(diff, key, value1)
		}
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

func isBothNested(value1, value2 interface{}) bool {
	_, ok1 := value1.(map[string]interface{})
	_, ok2 := value2.(map[string]interface{})

	return ok1 && ok2
}

func addNestedDiff(key string, value1, value2 interface{}, diff *[]Diff) {
	map1, _ := value1.(map[string]interface{}) // { "setting1": "Value 1", "setting2": 200 }
	map2, _ := value2.(map[string]interface{}) // { "setting1": "Value 1", "setting3": true }
	var nestedDiff []Diff
	nestedDiff = GenerateDiff(map1, map2, &nestedDiff)
	for _, value := range nestedDiff {
		newKey := key + "." + value.Key
		value.Key = newKey
		*diff = append(*diff, value)
	}
}
