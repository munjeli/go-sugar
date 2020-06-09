package sugar

// Common test data
type kitty struct {
	name string
	coat string
}

func testWithRunes(s string) []interface{} {
	rs := []rune(s)
	is := make([]interface{}, 0, len(rs))
	for _, r := range rs {
		is = append(is, r)
	}
	return is
}

var (
	emptySlice          = []interface{}{}
	mixedSlice          = []interface{}{"kitty", 1, kitty{"shirls", "spotted"}}
	stringSliceNoDups   = []interface{}{"kitty", "cat", "bat", "sate"}
	stringSliceWithDups = []interface{}{"kitten", "dog", "kitten"}
	emptyMap            = map[interface{}]interface{}{}
)
