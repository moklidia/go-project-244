package diff

import "sort"

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

func GenerateDiff(data1, data2 map[string]interface{}) []Diff {
	var result []Diff
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
			if value1 == value2 {
				AddUnchanged(&result, key, value1)
			} else {
				AddChanged(&result, key, value1, value2)
			}
		}
		if !ok1 {
			AddAdded(&result, key, value2)
		}

		if !ok2 {
			AddRemoved(&result, key, value1)
		}
	}

	sort.Slice(result, func(a, b int) bool {
		return result[a].Key < result[b].Key
	})

	return result
}

func AddUnchanged(result *[]Diff, key string, value interface{}) {
	newEntry := Diff{Type: Unchanged, Key: key, Value: value}
	*result = append(*result, newEntry)
}

func AddChanged(result *[]Diff, key string, value1, value2 interface{}) {
	newEntry := Diff{Type: Changed, Key: key, OldValue: value1, NewValue: value2}
	*result = append(*result, newEntry)
}

func AddRemoved(result *[]Diff, key string, value interface{}) {
	newEntry := Diff{Type: Removed, Key: key, Value: value}
	*result = append(*result, newEntry)
}

func AddAdded(result *[]Diff, key string, value interface{}) {
	newEntry := Diff{Type: Added, Key: key, Value: value}
	*result = append(*result, newEntry)
}
