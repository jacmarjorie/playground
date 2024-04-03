package structures

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var test1 = `8
push 3
push 7
push 20
peek
pop
push 0
push 4
pop`

var test2 = `8
push 3
push 7
push 8
peek
pop
peek
pop
pop`

func Test_pushPopStructs(t *testing.T) {

	tests := []struct {
		name    string
		dstruct pushPopStruct
		input   string
		want    string
	}{
		{
			name:    "Test1 - stack",
			dstruct: &stack{},
			input:   test1,
			want:    "20 [3 7 0]",
		},
		{
			name:    "Test2 - stack",
			dstruct: &stack{},
			input:   test2,
			want:    "8 7 []",
		},
		{
			name:    "Test1 - queue",
			dstruct: &queue{},
			input:   test1,
			want:    "3 [20 0 4]",
		},
		{
			name:    "Test2 - queue",
			dstruct: &queue{},
			input:   test2,
			want:    "3 7 []",
		},
	}

	for _, tt := range tests {
		got := tt.dstruct.evaluate(strings.NewReader(tt.input))
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("test %v failed want-/got+ %v", tt.name, diff)
		}
	}

}
