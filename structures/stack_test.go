package structures

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_stack(t *testing.T) {

	tests := []struct {
		name    string
		dstruct stack
		input   string
		want    string
	}{
		{
			name:    "Test1 - stack",
			dstruct: stack{},
			input:   test1,
			want:    "20 &[3 7 0]",
		},
		{
			name:    "Test2 - stack",
			dstruct: stack{},
			input:   test2,
			want:    "8 7 &[]",
		},
	}

	for _, tt := range tests {
		got := tt.dstruct.evaluate(strings.NewReader(tt.input))
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("test %v failed want-/got+ %v", tt.name, diff)
		}
	}

}
