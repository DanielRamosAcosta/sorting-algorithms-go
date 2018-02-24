package aedasort

import "sort"

// SelectionSort sorts the given sortable item using Selection Sort algorithm
// https://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(s sort.Interface) {
	for i := 0; i < s.Len(); i++ {
		indexMin := i

		for j := i + 1; j < s.Len(); j++ {
			if s.Less(j, indexMin) {
				indexMin = j
			}
		}
		s.Swap(i, indexMin)
	}
}
