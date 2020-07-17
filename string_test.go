package gosugar

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsPermutation(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsPermutation(tt.args.str1, tt.args.str2))
		})
	}
}
