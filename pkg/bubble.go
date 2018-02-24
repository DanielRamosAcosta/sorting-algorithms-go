package aedasort

import (
	"sort"
)

// BubbleSort sorts the given sortable item using Bubble Sort algorithm
// https://en.wikipedia.org/wiki/Bubble_sort
func BubbleSort(s sort.Interface) {
	change := true
	for i := 1; i < s.Len() && change; i++ {
		change = false
		for j := s.Len() - 1; j >= i; j-- {
			if s.Less(j, j-1) {
				s.Swap(j-1, j)
				change = true
			}
		}
	}
}
