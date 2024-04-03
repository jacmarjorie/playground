package structures

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_rotateLeftTillZero(t *testing.T) {
	tests := []struct {
		name string
		q    *queue
		want []any
	}{
		{
			name: "Test1 - zero at end",
			q:    &queue{eles: []any{1, 2, 3, 4, "A", "B", "C", 0}},
			want: []any{0, 1, 2, 3, 4, "A", "B", "C"},
		},
		{
			name: "Test2 - zero at beginning",
			q:    &queue{eles: []any{0, 1, 2, 3, 4, "A", "B", "C"}},
			want: []any{0, 1, 2, 3, 4, "A", "B", "C"},
		},
		{
			name: "Test2 - zero in middle",
			q:    &queue{eles: []any{1, 2, 3, 4, 0, "A", "B", "C"}},
			want: []any{0, "A", "B", "C", 1, 2, 3, 4},
		},
		{
			name: "Test3 - no zero",
			q:    &queue{eles: []any{1, 2, 3, 4, "A", "B", "C"}},
			want: []any{1, 2, 3, 4, "A", "B", "C"},
		},
	}

	for _, tt := range tests {
		rotateLeftTillZero(tt.q)
		if diff := cmp.Diff(tt.want, tt.q.elements()); diff != "" {
			t.Fatalf("test %v failed want-/got+ %v", tt.name, diff)
		}
	}
}

func generateRandomQueue(size int) *queue {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	zeroLoc := rand.Intn(size)
	q := &queue{}
	for i := 1; i < size; i++ {
		if i == zeroLoc {
			q.push(0)
			continue
		}
		q.push(i)
	}
	return q
}

func benchmarkRotateZero(q *queue, b *testing.B) {
	for n := 0; n < b.N; n++ {
		rotateLeftTillZero(q)
	}
}

// 46275768	        22.84 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero10(b *testing.B) { benchmarkRotateZero(generateRandomQueue(10), b) }

// 43963417	        23.59 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero100(b *testing.B) { benchmarkRotateZero(generateRandomQueue(100), b) }

// 53593687	        21.45 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero1000(b *testing.B) {
	benchmarkRotateZero(generateRandomQueue(1000), b)
}

// 53095393	        20.95 ns/op	       0 B/op	       0 allocs/op
func Benchmark_rotateLeftTillZero10000(b *testing.B) {
	benchmarkRotateZero(generateRandomQueue(10000), b)
}
