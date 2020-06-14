package go-sugar

// I mostly broke these out by key and value which may seem like a lot of redundant code
// but I am still thinking about how to support slice values and it's also perhaps clearer
// to me as a user to have that separation rather than passing a flag on key or value.
import "reflect"

// RemoveFromMapByKey returns a map where the `key: value` pair is removed
// as identified by the key. Probably only works with simple maps.
func RemoveFromMapByKey(m map[interface{}]interface{}, key interface{}) map[interface{}]interface{} {
	removed := map[interface{}]interface{}{}
	for k, v := range m {
		if reflect.DeepEqual(key, k) {
			continue
		} else {
			removed[k] = v
		}
	}
	return removed
}

// RemoveFromMapByValue does the same kind of test as above, but works with the
// value. Likewise, I don't think you can pass an array value...
func RemoveFromMapByValue(m map[interface{}]interface{}, val interface{}) map[interface{}]interface{} {
	removed := map[interface{}]interface{}{}
	for k, v := range m {
		if reflect.DeepEqual(val, v) {
			continue
		} else {
			removed[k] = v
		}
	}
	return removed
}

// FilterMapByKeyCondition will loop through the map and filter out the
// items that match the function's condition. Function must return bool.
func FilterMapByKeyCondition(m map[interface{}]interface{}, f func(i interface{}) bool) (map[interface{}]interface{}, map[interface{}]interface{}) {
	found := make(map[interface{}]interface{})
	filtered := make(map[interface{}]interface{})
	for k, v := range m {
		if f(k) {
			found[k] = v
		} else {
			filtered[k] = v
		}

	}
	return found, filtered
}

// FilterMapByValueCondition will loop through the map and filter out the
// items that match the function's condition. Still thinking about how to support slice values...
func FilterMapByValueCondition(m map[interface{}]interface{}, f func(i interface{}) bool) (map[interface{}]interface{}, map[interface{}]interface{}) {
	found := make(map[interface{}]interface{})
	filtered := make(map[interface{}]interface{})
	for k, v := range m {
		if f(v) {
			found[k] = v
		} else {
			filtered[k] = v
		}

	}
	return found, filtered
}

// SameMap is a recursive comparison of keys and values in both maps.
func SameMap(m1 map[interface{}]interface{}, m2 map[interface{}]interface{}) bool {
	return reflect.DeepEqual(m1, m2)
}
