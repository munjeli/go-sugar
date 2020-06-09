package sugar

import (
	"reflect"
	"testing"
)

func TestRemoveFromMapByKey(t *testing.T) {
    tests := []struct {
        desc string
        m map[interface{}]interface{}
        k interface{}
        want map[interface{}]interface{}
    } {
        {
            desc: "empty map",
            m: emptyMap,
            k: "kitty",
            want: emptyMap,
        },
    }
    for _, test := range tests {
        t.Run(test.desc, func(t *testing.T) {
            removed := RemoveFromMapByKey(test.m, test.k)
			if !reflect.DeepEqual(removed, test.want) {
				t.Errorf("returned wrong map: want: %v, got: %v", removed, test.want)
			}
        })
    }
}