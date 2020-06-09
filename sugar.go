package sugar

// RunesToInterfaces is just sugar so I can compare
// strings as a slice of characters.
func RunesToInterfaces(rs []rune) []interface{} {
	is := make([]interface{}, 0, len(rs))
	for _, r := range rs {
		is = append(is, r)
	}
	return is
}

