package sugar

import "fmt"

// HasDupes
func HasDupes(is []interface{}) bool {
	m := Uniq(is)
	if len(is) > len(m) {
		return true
	}
	return false
}

// Uniq sends back a deduped slice.
func Uniq(is []interface{}) []interface{} {
	m := map[interface{}]int{}
	for _, i := range is {
		m[i] = 0
	}
	keys := make([]interface{}, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Remove will remove a given item from an array.
func Remove(is []interface{}, i interface{}) []interface{} {
	removed := []interface{}{}
	for _, intfc := range is {
		if intfc == i {
			continue
		} else {
			removed = append(removed, intfc)
		}
	}
	return removed
}

// InSlice returns a bool if the value is in a slice.
func InSlice(is []interface{}, i interface{}) bool {
	for _, intfc := range is {
		if intfc == i {
			return true
		}
	}
	return false
}

// CountDups returns a map with the number of instances of an
// interface.
func CountDups(is []interface{}) map[interface{}]int {
	c := map[interface{}]int{}
	for _, i := range is {
		c[i] = c[i] +1
	}
	return c
}

// Reverse returns a slice reversed.
//func Reverse(is []interface{}) []interface{} {}

// Pop returns the first element of a slice, and the tail.
func Pop(is []interface{}) (interface{}, []interface{}) {
	if len(is) == 0 {
		return nil, []interface{}{}
	}
	return is[0], is[1:]
}

// ReplaceInSlice replaces the old value in the array with the new one.
func ReplaceInSlice(is []interface{}, old, new interface{}) ([]interface{}, error) {
	replaced := []interface{}{}
	if old == nil {
		return replaced, fmt.Errorf("old value cannot be nil")
	}
	for _, i := range is {
		if i == old {
			replaced = append(replaced, new)
		} else {
			replaced = append(replaced, i)
		}
	}
	return replaced, nil
}

// CountInSlice returns the number of times an interface is in the slice.
func CountInSlice(is []interface{}, countThis interface{}) int {
	mapWithCounts := CountDups(is)
	return mapWithCounts[countThis]
}

