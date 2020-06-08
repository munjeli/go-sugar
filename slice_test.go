package sugar

import (
	"testing"
)

func TestSliceHasDupes(t *testing.T) {
	tests := []struct {
		desc  string
		testi []interface{}
		want  bool
	}{
		{
			desc:  "empty slice",
			testi: emptySlice,
			want:  false,
		},
		{
			desc:  "string",
			testi: stringSliceWithDups,
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
			b := SliceHasDupes(test.testi)
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
			s:    emptySlice,
			i:    0,
			want: false,
		},
		{
			desc: "slice with kitten",
			s:    stringSliceWithDups,
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

func TestCountDupsInSlice(t *testing.T) {
	tests := []struct {
		desc    string
		is      []interface{}
		result  map[interface{}]int
		wantErr bool
	}{
		{
			desc:    "empty interface",
			is:      emptySlice,
			result:  make(map[interface{}]int),
			wantErr: false,
		},
		{
			desc: "string slice",
			is:   stringSliceWithDups,
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
			m := CountDupsInSlice(test.is)
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
			is:      emptySlice,
			old:     nil,
			new:     "kitty",
			wantErr: true,
		},
		{
			desc:    "slice of strings",
			is:      stringSliceNoDups,
			old:     "bat",
			new:     "cat",
			wantErr: false,
		},
		{
			desc:    "slice of things",
			is:      mixedSlice,
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
			is: emptySlice,
			i: "kitten",
			count: 0,
		},
		{
			desc: "test with irregular slice",
			is: mixedSlice,
			i: "kitty",
			count: 1,
		},
		{
			desc: "test with different cat",
			is: mixedSlice,
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

func TestRemoveFromSlice(t *testing.T) {
	tests := []struct {
		desc string
		is []interface{}
		i interface{}
	} {
		{
			desc: "empty slice",
			is: emptySlice,
			i: kitty{"shirls", "spotted"},
		},
		{
			desc: "weird slice",
			is: mixedSlice,
			i: kitty{"shirls", "spotted"},
		},
		{
			desc: "remove dupes",
			is: stringSliceWithDups,
			i: "kitty",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			removed := RemoveFromSlice(test.is, test.i)
			if InSlice(removed, test.i) {
				t.Errorf("failed to remove %v from slice", test.i)
			}
		})
	}
}

func TestPopSlice(t *testing.T) {
	tests := []struct {
		desc string
		is []interface{}
		expectedLen int
	} {
		{
			desc: "empty slice",
			is: emptySlice,
			expectedLen: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			_, tail := PopSlice(test.is)
			if len(tail) != test.expectedLen {
				t.Errorf("failed to pop slice: got len: %v, want: %v", len(tail), test.expectedLen)
			}
		})
	}
}

func TestReverseSlice(t *testing.T) {
    tests := []struct {
        desc string
        is []interface{}
        want []interface{}
    } {
        {
            desc: "empty slice",
            is: emptySlice,
            want: emptySlice,
        },
		{
			desc: "string slice",
			is: stringSliceNoDups,
			want: []interface{}{"sate", "bat", "cat", "kitty"} ,
		},
    }
    for _, test := range tests {
        t.Run(test.desc, func(t *testing.T) {
            reversed := ReverseSlice(test.is)
            for idx, i := range reversed {
            	if reversed[idx] != test.want[idx] {
            		t.Errorf("failed to reverse slice: want: %v, got: %v", test.want[idx], i)
				}
			}

        })
    }
}

func TestFilterSliceByCondition(t *testing.T) {
    tests := []struct {
        desc string
        is []interface{}
        f func(interface{}) bool
        filtered []interface{}
        targets []interface{}
    } {
		{
			desc:     "empty slice",
			f: func(i interface{}) bool { if i == "cat" { return true}; return false},
			is:       emptySlice,
			filtered: emptySlice,
			targets:  emptySlice,
		},
    }
    for _, test := range tests {
        t.Run(test.desc, func(t *testing.T) {
            targets, filtered := FilterSliceByCondition(test.is, test.f)
            // Check the empty slice is working as expected
			if len(test.is) == 0 {
				if len(targets) != 0 && len(filtered) != 0 {
					t.Errorf("failed on empty slice")
				}
			}
        })
    }
}