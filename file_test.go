package gosugar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToString(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileToString(tt.args.p)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
