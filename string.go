package sugar

import "reflect"

// IsPermutation tells if str2 a permutation of str1.
// We compare maps of the string where the value is the number
// of times a key appears. The same number of same characters
// is a permutation.
func IsPermutation(str1, str2 string) bool {
	if str1 == str2 {
		return true
	}
	r1 := []rune(str1)
	r2 := []rune(str2)
	runeMap1 := CountDupsInSlice(RunesToInterfaces(r1))
	runeMap2 := CountDupsInSlice(RunesToInterfaces(r2))
	sameMap := reflect.DeepEqual(runeMap1, runeMap2)
	if !sameMap {
		return false
	}
	return true
}
