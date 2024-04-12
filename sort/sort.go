package sort

func insertionSort(lst []int) []int {
	for i := 0; i < len(lst); i++ {
		cur := i
		for cur > 0 && lst[cur] < lst[cur-1] {
			tmp := lst[cur-1]
			lst[cur-1] = lst[cur]
			lst[cur] = tmp
			cur -= 1
		}
	}
	return lst
}

func selectionSort(lst []int) []int {
	for j := 0; j < len(lst); j++ {
		min_index := j
		for i := j + 1; i < len(lst); i++ {
			if lst[i] < lst[min_index] {
				min_index = i
			}
		}
		tmp := lst[j]
		lst[j] = lst[min_index]
		lst[min_index] = tmp
	}
	return lst
}

func bubbleSort(lst []int) []int {
	swap := false
	for i := 0; i < len(lst); i++ {
		for j := i + 1; j < len(lst); j++ {
			if lst[j] < lst[i] {
				swap = true
				tmp := lst[i]
				lst[i] = lst[j]
				lst[j] = tmp
			}
		}
		if !swap {
			return lst
		}
	}
	return lst
}

func mergeSort(lst []int) []int {
	// empty or len 1 list is already sorted
	if len(lst) <= 1 {
		return lst
	}

	// find the mid point
	mid := len(lst) / 2
	// split and sort
	l, r := mergeSort(lst[:mid]), mergeSort(lst[mid:])
	result := make([]int, 0, len(lst))
	lPtr, rPtr := 0, 0

	// while both not at midpoint
	for lPtr < mid || rPtr < len(lst)-mid {
		if lPtr == mid {
			result = append(result, r[rPtr])
			rPtr++
		} else if rPtr == len(lst)-mid {
			result = append(result, l[lPtr])
			lPtr++
		} else if l[lPtr] <= r[rPtr] {
			result = append(result, l[lPtr])
			lPtr++
		} else {
			result = append(result, r[rPtr])
			rPtr++
		}
	}
	return result
}

// todo make pivot the mid point
func quickSortRecur(lst []int, start, end int) {
	if end-start <= 1 {
		return
	}
	pivot := lst[end-1]
	startPtr, endPtr := start, end-1

	for startPtr < endPtr {
		for lst[startPtr] < pivot && startPtr < endPtr {
			startPtr++
		}

		for lst[endPtr] >= pivot && startPtr < endPtr {
			endPtr--
		}

		lst[startPtr], lst[endPtr] = lst[endPtr], lst[startPtr]
	}

	lst[startPtr], lst[end-1] = lst[end-1], lst[startPtr]

	quickSortRecur(lst, start, startPtr)
	quickSortRecur(lst, startPtr+1, end)
}

func quickSort(lst []int) {
	quickSortRecur(lst, 0, len(lst))
}
