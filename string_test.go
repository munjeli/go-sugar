package gosugar

import "testing"

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		desc string
		str1 string
		str2 string
		want bool
	}{
		{
			desc: "empty strings",
			str1: "",
			str2: "",
			want: true,
		},
		{
			desc: "same string",
			str1: "shuffle",
			str2: "shuffle",
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			b := IsPermutation(test.str1, test.str2)
			if b != test.want {
				t.Errorf("want: %v, got: %v", test.want, b)
			}
		})
	}
}
