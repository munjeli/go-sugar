package gosugar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mapWithStringKeys = map[interface{}]interface{}{
		"kitty": "Shirley",
		"puppy": "Neptune",
		"rat":   "Squeaky",
	}
	mapWithIntKeys = map[interface{}]interface{}{
		1: "Shirley",
		2: "Neptune",
		3: "Squeaky",
	}
	mapWithIntVals = map[interface{}]interface{}{
		"Shirley": 1,
		"Neptune": 2,
		"Squeaky": 3,
	}
)

func isHello(i interface{}) bool {
	if i == "Hello" {
		return true
	}
	return false
}

func TestRemoveFromMapByKey(t *testing.T) {
	type args struct {
		m   map[interface{}]interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]interface{}
	}{
		{
			name: "empty map",
			args: args{
				m:   map[interface{}]interface{}{},
				key: "kitty",
			},
			want: map[interface{}]interface{}{},
		},
		{
			name: "map with string keys as interfaces",
			args: args{
				m:   mapWithStringKeys,
				key: "kitty",
			},
			want: map[interface{}]interface{}{
				"puppy": "Neptune",
				"rat":   "Squeaky",
			},
		},
		{
			name: "map with int keys as interfaces",
			args: args{
				m:   mapWithIntKeys,
				key: 1,
			},
			want: map[interface{}]interface{}{
				2: "Neptune",
				3: "Squeaky",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveFromMapByKey(tt.args.m, tt.args.key))
		})
	}
}

func TestRemoveFromMapByValue(t *testing.T) {
	type args struct {
		m   map[interface{}]interface{}
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]interface{}
	}{
		{
			name: "empty map",
			args: args{
				m:   map[interface{}]interface{}{},
				val: 1,
			},
			want: map[interface{}]interface{}{},
		},
		{
			name: "map with string vals as interfaces",
			args: args{
				m:   mapWithStringKeys,
				val: "Shirley",
			},
			want: map[interface{}]interface{}{
				"puppy": "Neptune",
				"rat":   "Squeaky",
			},
		},
		{
			name: "map with int vals as interfaces",
			args: args{
				m:   mapWithIntVals,
				val: 1,
			},
			want: map[interface{}]interface{}{
				"Neptune": 2,
				"Squeaky": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveFromMapByValue(tt.args.m, tt.args.val))
		})
	}
}

func TestFilterMapByKeyCondition(t *testing.T) {
	type args struct {
		m map[interface{}]interface{}
		f func(i interface{}) bool
	}
	tests := []struct {
		name  string
		args  args
		want  map[interface{}]interface{}
		want1 map[interface{}]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FilterMapByKeyCondition(tt.args.m, tt.args.f)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestFilterMapByValueCondition(t *testing.T) {
	type args struct {
		m map[interface{}]interface{}
		f func(i interface{}) bool
	}
	tests := []struct {
		name  string
		args  args
		want  map[interface{}]interface{}
		want1 map[interface{}]interface{}
	}{
		{
			name: "empty map",
			args: args{
				m: emptyMap,
				f: isHello,
			},
			want:  emptyMap,
			want1: emptyMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FilterMapByValueCondition(tt.args.m, tt.args.f)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestSameMap(t *testing.T) {
	type args struct {
		m1 map[interface{}]interface{}
		m2 map[interface{}]interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty map",
			args: args{
				m1: map[interface{}]interface{}{},
				m2: map[interface{}]interface{}{},
			},
			want: true,
		},
		{
			name: "same map",
			args: args{
				m1: mapWithStringKeys,
				m2: mapWithStringKeys,
			},
			want: true,
		},
		{
			name: "different map",
			args: args{
				m1: mapWithStringKeys,
				m2: mapWithIntVals,
			},
			want: false,
		},
		{
			name: "first map empty",
			args: args{
				m1: map[interface{}]interface{}{},
				m2: mapWithIntVals,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SameMap(tt.args.m1, tt.args.m2))
		})
	}
}

func TestSortMapByStringKey(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "empty map",
			args: args{
				m: map[string]interface{}{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SortMapByStringKey(tt.args.m))
		})
	}
}

func TestSortMapByIntKey(t *testing.T) {
	type args struct {
		m map[int]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[int]interface{}
	}{
		{
			name: "empty map",
			args: args{
				m: map[int]interface{}{},
			},
			want: map[int]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SortMapByIntKey(tt.args.m))
		})
	}
}

func TestSortMapByStringValue(t *testing.T) {
	type args struct {
		m map[interface{}]string
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]string
	}{
		{
			name: "empty map",
			args: args{
				m: map[interface{}]string{},
			},
			want: map[interface{}]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SortMapByStringValue(tt.args.m))
		})
	}
}
