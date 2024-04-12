package binary

import "log"

func feasible[T comparable](i T, target T) bool {
	return i == target
}

func binarySearchFeasible[T comparable](lst []T, f func(T) bool) int {
	l, r := 0, len(lst)-1
	cur := -1
	for l <= r {
		n := l + (r-l)/2
		log.Printf("%v", n)
		if f(lst[n]) {
			cur = n
			r = n - 1
		} else {
			l = n + 1
		}
	}
	return cur

}

func binarySearch(lst []int, target int) int {
	l, r := 0, len(lst)-1
	for l <= r {
		n := l + (r-l)/2
		if lst[n] == target {
			return n
		} else if lst[n] < target {
			l = n + 1
		} else if lst[n] > target {
			r = n - 1
		}
	}
	return -1
}

func binarySearchBool(lst []bool) int {
	l, r := 0, len(lst)-1
	cur := -1
	for l <= r {
		n := l + (r-l)/2
		if lst[n] { // ... T ...
			cur = n
			r = n - 1
		} else {
			l = n + 1
		}
	}
	return cur
}

func binarySearchFindInDups(lst []int, target int) int {
	l, r := 0, len(lst)-1
	cur := -1
	for l <= r {
		n := l + (r-l)/2
		if lst[n] == target {
			cur = n
			r = n - 1
		} else if lst[n] > target {
			r = n - 1
		} else {
			l = n + 1
		}
	}
	return cur

}
