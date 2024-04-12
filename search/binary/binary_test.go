package binary

import (
	"sort"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		target   int
		expected int
	}{
		{
			name:     "Test1",
			lst:      []int{1, 3, 5, 7, 8},
			target:   5,
			expected: 2,
		},
		{
			name:     "Test2",
			lst:      []int{1, 2, 3, 4, 5, 6, 7},
			target:   0,
			expected: -1,
		},
		{
			name:     "Test3",
			lst:      []int{2, 8, 89, 120, 1000},
			target:   120,
			expected: 3,
		},
		{
			name:     "Test4",
			lst:      []int{10, 20},
			target:   20,
			expected: 1,
		},
		{
			name:     "Test5",
			lst:      []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Test6",
			lst:      []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "Test7",
			lst:      []int{1, 2, 3, 4, 5},
			target:   10,
			expected: -1,
		},
	}
	for _, tt := range tests {
		sort.Ints(tt.lst)
		got := binarySearch(tt.lst, tt.target)
		var intF = func(i int) bool { return feasible(i, tt.target) }
		gotFeasible := binarySearchFeasible(tt.lst, intF)
		if got != tt.expected || gotFeasible != tt.expected {
			t.Fatalf("test %v failed: want=%v got=%v gotFeasible=%v", tt.name, tt.expected, got, gotFeasible)
		}

	}
}

func Test_binarySearchBool(t *testing.T) {
	tests := []struct {
		name     string
		lst      []bool
		expected int
	}{
		{
			name:     "Test1",
			lst:      []bool{false, false, true, true, true},
			expected: 2,
		},
		{
			name:     "Test2",
			lst:      []bool{true},
			expected: 0,
		},
		{
			name:     "Test3",
			lst:      []bool{false, false, false},
			expected: -1,
		},
		{
			name:     "Test4",
			lst:      []bool{true, true, true, true},
			expected: 0,
		},
		{
			name:     "Test5",
			lst:      []bool{false, true},
			expected: 1,
		},
		{
			name:     "Test6",
			lst:      []bool{false, false, false, false, false, false, false, false, true},
			expected: 8,
		},
		{
			name:     "Test7",
			lst:      []bool{},
			expected: -1,
		},
	}
	for _, tt := range tests {
		got := binarySearchBool(tt.lst)
		var boolF = func(i bool) bool { return feasible(i, true) }
		gotFeasible := binarySearchFeasible(tt.lst, boolF)
		if got != tt.expected || gotFeasible != tt.expected {
			t.Fatalf("test %v failed: want=%v got=%v gotFeasible = %v", tt.name, tt.expected, got, gotFeasible)
		}

	}
}

func Test_firstElementNotSmallerThanTarget(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		target   int
		expected int
	}{
		{
			name:     "Test 1",
			lst:      []int{1, 3, 3, 5, 8, 8, 10},
			target:   2,
			expected: 1,
		},
		{
			name:     "Test 2",
			lst:      []int{2, 3, 5, 7, 11, 13, 17, 19},
			target:   6,
			expected: 3,
		},
	}

	for _, tt := range tests {
		var testFeasible = func(i int) bool {
			return i >= tt.target
		}
		got := binarySearchFeasible(tt.lst, testFeasible)
		if got != tt.expected {
			t.Fatalf("test %v failed: want=%v got=%v", tt.name, tt.expected, got)
		}
	}
}

func Test_findElementInSortedWithDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		lst      []int
		target   int
		expected int
	}{
		{
			name:     "Test 1",
			lst:      []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100},
			target:   3,
			expected: 1,
		},
		{
			name:     "Test 2",
			lst:      []int{2, 3, 5, 7, 11, 13, 17, 19},
			target:   6,
			expected: -1,
		},
		{
			name:     "Test 3",
			lst:      []int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100},
			target:   3,
			expected: 1,
		},
		{
			name:     "Test 4",
			lst:      []int{1, 3, 3, 3, 6, 6, 10, 10, 10, 100},
			target:   3,
			expected: 1,
		},
	}

	for _, tt := range tests {
		got := binarySearchFindInDups(tt.lst, tt.target)
		if got != tt.expected {
			t.Fatalf("test %v failed: want=%v got=%v", tt.name, tt.expected, got)
		}
	}
}
