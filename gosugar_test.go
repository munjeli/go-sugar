package gosugar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Common test data
type kitty struct {
	name string
	coat string
}

var (
	emptySlice          = []interface{}{}
	mixedSlice          = []interface{}{"kitty", 1, kitty{"shirls", "spotted"}}
	stringSliceNoDups   = []interface{}{"kitty", "cat", "bat", "sate"}
	stringSliceWithDups = []interface{}{"kitten", "dog", "kitten"}
	emptyMap            = map[interface{}]interface{}{}
)

func testWithRunes(s string) []interface{} {
	rs := []rune(s)
	is := make([]interface{}, 0, len(rs))
	for _, r := range rs {
		is = append(is, r)
	}
	return is
}

func TestRunesToInterfaces(t *testing.T) {
	type args struct {
		rs []rune
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RunesToInterfaces(tt.args.rs))
		})
	}
}
