package structures

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_queue(t *testing.T) {

	tests := []struct {
		name    string
		dstruct queue
		input   string
		want    string
	}{
		{
			name:    "Test1 - queue",
			dstruct: queue{},
			input:   test1,
			want:    "3 &[20 0 4]",
		},
		{
			name:    "Test2 - queue",
			dstruct: queue{},
			input:   test2,
			want:    "3 7 &[]",
		},
	}

	for _, tt := range tests {
		got := tt.dstruct.evaluate(strings.NewReader(tt.input))
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("test %v failed want-/got+ %v", tt.name, diff)
		}
	}

}

func Test_rotateLeftTillZero(t *testing.T) {
	tests := []struct {
		name string
		q    queue
		want queue
	}{
		{
			name: "Test1 - zero at end",
			q:    queue{1, 2, 3, 4, "A", "B", "C", 0},
			want: queue{0, 1, 2, 3, 4, "A", "B", "C"},
		},
		{
			name: "Test2 - zero at beginning",
			q:    queue{0, 1, 2, 3, 4, "A", "B", "C"},
			want: queue{0, 1, 2, 3, 4, "A", "B", "C"},
		},
		{
			name: "Test2 - zero in middle",
			q:    queue{1, 2, 3, 4, 0, "A", "B", "C"},
			want: queue{0, "A", "B", "C", 1, 2, 3, 4},
		},
		{
			name: "Test3 - no zero",
			q:    queue{1, 2, 3, 4, "A", "B", "C"},
			want: queue{1, 2, 3, 4, "A", "B", "C"},
		},
		{
			name: "Test4 - empty list",
			q:    queue{},
			want: queue{},
		},
	}

	for _, tt := range tests {
		rotateLeftTillZero(&tt.q)
		if diff := cmp.Diff(tt.want, tt.q); diff != "" {
			t.Fatalf("test %v failed want-/got+ %v", tt.name, diff)
		}
	}
}

func benchmarkRotateZero(q *queue, b *testing.B) {
	for n := 0; n < b.N; n++ {
		rotateLeftTillZero(q)
	}
}

// 46275768	        22.84 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero10(b *testing.B) { benchmarkRotateZero(GenerateRandomQueue(10), b) }

// 43963417	        23.59 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero100(b *testing.B) { benchmarkRotateZero(GenerateRandomQueue(100), b) }

// 53593687	        21.45 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero1000(b *testing.B) {
	benchmarkRotateZero(GenerateRandomQueue(1000), b)
}

// 53095393	        20.95 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero10000(b *testing.B) {
	benchmarkRotateZero(GenerateRandomQueue(10000), b)
}
