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
		{
			name: "empty strings",
			args: args{
				str1: "",
				str2: "",
			},
			want: true,
		},
		{
			name: "is permutation",
			args: args{
				str1: "caats",
				str2: "tasac",
			},
			want: true,
		},
		{
			name: "not permutation",
			args: args{
				str1: "puppies",
				str2: "kitties",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsPermutation(tt.args.str1, tt.args.str2))
		})
	}
}
