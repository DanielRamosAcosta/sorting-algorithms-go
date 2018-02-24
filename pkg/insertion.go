package aedasort

import (
	"sort"
)

// InsertionSort sorts the given sortable item using Insertion Sort algorithm
func InsertionSort(s sort.Interface) {
	for i, j := 1, 1; i < s.Len(); i, j = i+1, i+1 {
		for (j > 0) && s.Less(j, j-1) {
			s.Swap(j, j-1)
			j--
		}
	}
}
