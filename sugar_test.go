package main

import (
	"fmt"
	"testing"
)

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

func TestHasDupes(t *testing.T) {
	tests := []struct {
		desc  string
		testi []interface{}
		want  bool
	}{
		{
			desc:  "empty slice",
			testi: []interface{}{},
			want:  false,
		},
		{
			desc:  "string",
			testi: []interface{}{"foo", "kitten", "lala", "foo"},
			want:  true,
		},
		{
			desc:  "runes",
			testi: testWithRunes("abcdefghijklmnop"),
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			b := HasDupes(test.testi)
			if b != test.want {
				t.Errorf("want: %v, got: %v", test.want, b)
			}
		})
	}
}

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

func TestInSlice(t *testing.T) {
	tests := []struct {
		desc string
		s    []interface{}
		i    interface{}
		want bool
	}{
		{
			desc: "empty slice",
			s:    []interface{}{},
			i:    0,
			want: false,
		},
		{
			desc: "slice with kitten",
			s:    []interface{}{"kitten", "cat", "dog"},
			i:    "kitten",
			want: true,
		},
		{
			desc: "slice without kitten",
			s:    []interface{}{"platypus", "cat", "dog"},
			i:    "kitten",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if InSlice(test.s, test.i) != test.want {
				t.Errorf("wrong bool: want: %v", test.want)
			}
		})
	}
}

func TestCountDups(t *testing.T) {
	tests := []struct {
		desc    string
		is      []interface{}
		result  map[interface{}]int
		wantErr bool
	}{
		{
			desc:    "empty interface",
			is:      []interface{}{},
			result:  make(map[interface{}]int),
			wantErr: false,
		},
		{
			desc: "string slice",
			is:   []interface{}{"kitten", "dog", "kitten"},
			result: map[interface{}]int{
				"kitten": 2,
				"dog":    1,
			},
			wantErr: false,
		},
		{
			desc: "struct slice",
			is:   []interface{}{kitty{name: "shirls", coat: "spotted"}, kitty{name: "shirls", coat: "spotted"}},
			result: map[interface{}]int{
				kitty{"shirls", "spotted"}: 2,
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			m := CountDups(test.is)
			for k, v := range m {
				if val, ok := test.result[k]; ok {
					if val != v && !test.wantErr {
						t.Errorf("incorrect value: want: %v, got: %v", test.result[k], v)
					}
				} else {
					if !test.wantErr {
						t.Errorf("missing key: %v", k)
					}
				}

			}
		})
	}
}

func TestReplaceInSlice(t *testing.T) {
	tests := []struct {
		desc    string
		is      []interface{}
		old     interface{}
		new     interface{}
		wantErr bool
	}{
		{
			desc:    "empty slice",
			is:      []interface{}{},
			old:     nil,
			new:     "kitty",
			wantErr: true,
		},
		{
			desc:    "slice of strings",
			is:      []interface{}{"kitty", "cat", "bat", "sate"},
			old:     "bat",
			new:     "cat",
			wantErr: false,
		},
		{
			desc:    "slice of things",
			is:      []interface{}{"kitty", 1, kitty{"shirls", "spotted"}},
			old:     1,
			new:     "cat",
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			newArr, err := ReplaceInSlice(test.is, test.old, test.new)
			if err != nil && !test.wantErr {
				t.Errorf("unexpected test failure: %v", err)
			}
			if InSlice(newArr, test.old) && test.old != nil {
				t.Errorf("failed to replace values in slice")
			}
			if !InSlice(newArr, test.new) && test.old != nil {
				t.Errorf("failed to remove old value value: %v", test.old)
			}
			if len(test.is) != len(newArr) {
				t.Errorf("something wrong with array length")
			}

		})

	}
}

func TestCountInSlice(t *testing.T) {
    tests := []struct {
        desc string
        is []interface{}
        i interface{}
        count int
    } {
        {
            desc: "test with empty slice",
            is: []interface{}{},
            i: "kitten",
            count: 0,
        },
		{
			desc: "test with irregular slice",
			is: []interface{}{"kitty", 1, kitty{"shirls", "spotted"}, "kitty"},
			i: "kitty",
			count: 2,
		},
		{
			desc: "test with different cat",
			is: []interface{}{"kitty", 1, kitty{"shirls", "spotted"}, "kitty"},
			i: kitty{"fish", "marmalade"},
			count: 0,
		},
    }
    for _, test := range tests {
        t.Run(test.desc, func(t *testing.T) {
        	c := CountInSlice(test.is, test.i)
            if c != test.count {
                t.Errorf("wrong count: want: %v, got: %v", test.count, c)
            }
        })
    }
}

func TestRemove(t *testing.T) {
    tests := []struct {
        desc string
        is []interface{}
        i interface{}
    } {
        {
            desc: "empty slice",
            is: []interface{}{},
            i: kitty{"shirls", "spotted"},
        },
		{
			desc: "weird slice",
			is: []interface{}{"kitty", 1, kitty{"shirls", "spotted"}, "kitty"},
			i: kitty{"shirls", "spotted"},
		},
		{
			desc: "remove dupes",
			is: []interface{}{"kitty", 1, kitty{"shirls", "spotted"}, "kitty"},
			i: "kitty",
		},
    }
    for _, test := range tests {
        t.Run(test.desc, func(t *testing.T) {
            removed := Remove(test.is, test.i)
            fmt.Println(test.i)
            if InSlice(removed, test.i) {
            	t.Errorf("failed to remove %v from slice", test.i)
			}
        })
    }
}

func TestPop(t *testing.T) {
    tests := []struct {
        desc string
        is []interface{}
        expectedLen int
    } {
        {
            desc: "empty slice",
            is: []interface{}{},
            expectedLen: 0,
        },
    }
    for _, test := range tests {
        t.Run(test.desc, func(t *testing.T) {
           	_, tail := Pop(test.is)
            if len(tail) != test.expectedLen {
                t.Errorf("failed to pop slice: got len: %v, want: %v", len(tail), test.expectedLen)
            }
        })
    }
}