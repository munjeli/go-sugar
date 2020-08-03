package gosugar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceHasDupes(t *testing.T) {
	type args struct {
		is []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty slice",
			args: args{is: emptySlice},
			want: false,
		},
		{
			name: "string",
			args: args{is: stringSliceWithDups},
			want: true,
		},
		{
			name: "runes",
			args: args{is: testWithRunes("abcdefghijklmnop")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SliceHasDupes(tt.args.is))
		})
	}
}

func TestUniqSlice(t *testing.T) {
	type args struct {
		is []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "empty slice",
			args: args{
				is: []interface{}{},
			},
			want: []interface{}{},
		},
		{
			name: "with some dupes",
			args: args{
				is: stringSliceWithDups,
			},
			want: []interface{}{"kitten", "dog"},
		},
		{
			name: "no dups",
			args: args{
				is: stringSliceNoDups,
			},
			want: stringSliceNoDups,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, UniqSlice(tt.args.is))
		})
	}
}

func TestRemoveFromSlice(t *testing.T) {
	type args struct {
		is []interface{}
		i  interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "empty slice",
			args: args{
				is: emptySlice,
				i:  kitty{"shirls", "spotted"},
			},
		},
		{
			name: "weird slice",
			args: args{
				is: mixedSlice,
				i:  kitty{"shirls", "spotted"},
			},
		},
		{
			name: "remove dupes",
			args: args{
				is: stringSliceWithDups,
				i:  "kitty",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removed := RemoveFromSlice(tt.args.is, tt.args.i)
			if InSlice(removed, tt.args.i) {
				t.Errorf("failed to remove %v from slice", tt.args.i)
			}
		})
	}
}

func TestInSlice(t *testing.T) {
	type args struct {
		is []interface{}
		i  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty slice",
			args: args{
				is: emptySlice,
				i:  0,
			},
			want: false,
		},
		{
			name: "slice with kitten",
			args: args{
				is: stringSliceWithDups,
				i:  "kitten",
			},
			want: true,
		},
		{
			name: "slice without kitten",
			args: args{
				is: []interface{}{"platypus", "cat", "dog"},
				i:  "kitten",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, InSlice(tt.args.is, tt.args.i))
		})
	}
}

func TestCountDupsInSlice(t *testing.T) {
	type args struct {
		is []interface{}
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CountDupsInSlice(tt.args.is))
		})
	}
}

func TestReverseSlice(t *testing.T) {
	type args struct {
		is []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "empty slice",
			args: args{
				is: emptySlice,
			},
			want: emptySlice,
		},
		{
			name: "string slice",
			args: args{
				is: stringSliceNoDups,
			},
			want: []interface{}{"sate", "bat", "cat", "kitty"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rvd := ReverseSlice(tt.args.is)
			assert.Equal(t, tt.want, rvd)
		})
	}
}

func TestFilterSliceByCondition(t *testing.T) {
	type args struct {
		is []interface{}
		f  func(i interface{}) bool
	}
	tests := []struct {
		name         string
		args         args
		wantTargets  []interface{}
		wantFiltered []interface{}
	}{
		{
			name: "empty interface",
			args: args{
				is: []interface{}{},
				f:  isHello,
			},
			wantTargets:  emptySlice,
			wantFiltered: emptySlice,
		},
		{
			name: "with Hello",
			args: args{
				is: stringSliceHello,
				f:  isHello,
			},
			wantTargets:  []interface{}{"Hello"},
			wantFiltered: []interface{}{"kitten", "here kitty kitty"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTargets, gotFiltered := FilterSliceByCondition(tt.args.is, tt.args.f)
			assert.Equal(t, tt.wantTargets, gotTargets)
			assert.Equal(t, tt.wantFiltered, gotFiltered)
		})
	}
}

func TestPopSlice(t *testing.T) {
	type args struct {
		is []interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 []interface{}
	}{
		{
			name: "string slice",
			args: args{
				is: stringSliceNoDups,
			},
			want: "sate",
			want1: []interface{}{
				"bat",
				"cat",
				"kitty",
			},
		},
		{
			name: "empty slice",
			args: args{
				is: emptySlice,
			},
			want:  nil,
			want1: emptySlice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PopSlice(tt.args.is)
			assert.Equal(t, tt.want, got)
			assert.ElementsMatch(t, tt.want1, got1)
		})
	}
}

func TestReplaceInSlice(t *testing.T) {
	type args struct {
		is  []interface{}
		old interface{}
		new interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "empty slice",
			args: args{
				is:  []interface{}{},
				old: "kitty",
				new: "puppy",
			},
			want:    []interface{}{},
			wantErr: false,
		},
		{
			name: "string slice",
			args: args{
				is:  stringSliceNoDups,
				old: "kitty",
				new: "puppy",
			},
			want:    []interface{}{"puppy", "bat", "cat", "sate"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReplaceInSlice(tt.args.is, tt.args.old, tt.args.new)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestCountInSlice(t *testing.T) {
	type args struct {
		is []interface{}
		i  interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with empty slice",
			args: args{
				is: emptySlice,
				i:  "kitten",
			},
			want: 0,
		},
		{
			name: "test with irregular slice",
			args: args{
				is: mixedSlice,
				i:  "kitty",
			},
			want: 1,
		},
		{
			name: "test with different cat",
			args: args{
				is: mixedSlice,
				i:  kitty{"fish", "marmalade"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CountInSlice(tt.args.is, tt.args.i))
		})
	}
}
