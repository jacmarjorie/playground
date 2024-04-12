package sort

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insertionSort(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		expected []int
	}{
		{
			name:     "Test1",
			lst:      []int{5, 3, 1, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test2",
			lst:      []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Test3",
			lst:      []int{8466, 1024, 7744, 4668, 2011, 7744, 6861, 8964, 1100},
			expected: []int{1024, 1100, 2011, 4668, 6861, 7744, 7744, 8466, 8964},
		},
	}
	for _, tt := range tests {
		got := insertionSort(tt.lst)
		if diff := cmp.Diff(got, tt.expected); diff != "" {
			t.Fatalf("test %v failed, want-/got+: %v %v", tt.name, diff, got)
		}
	}
}

func Test_selectionSort(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		expected []int
	}{
		{
			name:     "Test1",
			lst:      []int{5, 3, 1, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test2",
			lst:      []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Test3",
			lst:      []int{8466, 1024, 7744, 4668, 2011, 7744, 6861, 8964, 1100},
			expected: []int{1024, 1100, 2011, 4668, 6861, 7744, 7744, 8466, 8964},
		},
	}
	for _, tt := range tests {
		got := selectionSort(tt.lst)
		if diff := cmp.Diff(got, tt.expected); diff != "" {
			t.Fatalf("test %v failed, want-/got+: %v %v", tt.name, diff, got)
		}
	}
}

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		expected []int
	}{
		{
			name:     "Test1",
			lst:      []int{5, 3, 1, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test2",
			lst:      []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Test3",
			lst:      []int{8466, 1024, 7744, 4668, 2011, 7744, 6861, 8964, 1100},
			expected: []int{1024, 1100, 2011, 4668, 6861, 7744, 7744, 8466, 8964},
		},
	}
	for _, tt := range tests {
		got := bubbleSort(tt.lst)
		if diff := cmp.Diff(got, tt.expected); diff != "" {
			t.Fatalf("test %v failed, want-/got+: %v %v", tt.name, diff, got)
		}
	}
}

func Test_mergeSort(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		expected []int
	}{
		{
			name:     "Test1",
			lst:      []int{5, 3, 1, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test2",
			lst:      []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Test3",
			lst:      []int{8466, 1024, 7744, 4668, 2011, 7744, 6861, 8964, 1100},
			expected: []int{1024, 1100, 2011, 4668, 6861, 7744, 7744, 8466, 8964},
		},
	}
	for _, tt := range tests {
		got := mergeSort(tt.lst)
		if diff := cmp.Diff(got, tt.expected); diff != "" {
			t.Fatalf("test %v failed, want-/got+: %v %v", tt.name, diff, got)
		}
	}
}

func Test_quickSort(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		expected []int
	}{
		{
			name:     "Test1",
			lst:      []int{5, 3, 1, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test2",
			lst:      []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Test3",
			lst:      []int{8466, 1024, 7744, 4668, 2011, 7744, 6861, 8964, 1100},
			expected: []int{1024, 1100, 2011, 4668, 6861, 7744, 7744, 8466, 8964},
		},
	}
	for _, tt := range tests {
		quickSort(tt.lst)
		if diff := cmp.Diff(tt.lst, tt.expected); diff != "" {
			t.Fatalf("test %v failed, want-/got+: %v %v", tt.name, diff, tt.lst)
		}
	}
}
