package gosugar

// SliceHasDupes returns a bool if a slice has duplicate interfaces.
func SliceHasDupes(is []interface{}) bool {
	m := UniqSlice(is)
	if len(is) > len(m) {
		return true
	}
	return false
}

// UniqSlice sends back a deduped slice.
func UniqSlice(is []interface{}) []interface{} {
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

// RemoveFromSlice will remove a given item from a slice.
func RemoveFromSlice(is []interface{}, i interface{}) []interface{} {
	var removed []interface{}
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

// CountDupsInSlice returns a map with the number of instances of an
// interface.
func CountDupsInSlice(is []interface{}) map[interface{}]int {
	c := map[interface{}]int{}
	for _, i := range is {
		c[i] = c[i] + 1
	}
	return c
}

// ReverseSlice returns a slice reversed.
func ReverseSlice(is []interface{}) []interface{} {
	for i, nexti := 0, len(is)-1; i < nexti; i, nexti = i+1, nexti-1 {
		is[i], is[nexti] = is[nexti], is[i]
	}
	return is
}

// FilterSliceByCondition takes a slice and a function that takes interface an returns a bool. This is
// more flexible than writing lots of conditions naturally, but you have to craft a function.
// example:
// f := func(i interface{}) bool { if i == "cat"; return true }
func FilterSliceByCondition(is []interface{}, f func(i interface{}) bool) (targets []interface{}, filtered []interface{}) {
	for _, i := range is {
		if f(i) {
			targets = append(targets, i)
		} else {
			filtered = append(filtered, i)
		}
	}
	return targets, filtered
}

// PopSlice returns the first element of a slice, and the tail.
// Don't forget to overwrite the var if you're using this in a loop.
func PopSlice(is []interface{}) (interface{}, []interface{}) {
	if len(is) == 0 {
		return nil, []interface{}{}
	}
	return is[0], is[1:]
}

// ReplaceInSlice replaces the old value in the array with the new one.
func ReplaceInSlice(is []interface{}, old, new interface{}) []interface{} {
	replaced := []interface{}{}
	for _, i := range is {
		if i == old {
			replaced = append(replaced, new)
		} else {
			replaced = append(replaced, i)
		}
	}
	return replaced
}

// CountInSlice returns the number of times an interface is in the slice.
func CountInSlice(is []interface{}, i interface{}) int {
	mapWithCounts := CountDupsInSlice(is)
	return mapWithCounts[i]
}
