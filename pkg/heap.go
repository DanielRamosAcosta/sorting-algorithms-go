package aedasort

import (
	"sort"
)

// HeapSort sorts the given sortable item using Heap Sort algorithm
// https://en.wikipedia.org/wiki/Heapsort
func HeapSort(s sort.Interface) {
	for i := s.Len()/2 - 1; i >= 0; i-- {
		shiftDown(s, i, s.Len())
	}

	for i := s.Len() - 1; i > 0; i-- {
		s.Swap(0, i)
		shiftDown(s, 0, i-1)
	}
}

func shiftDown(s sort.Interface, i int, n int) {
	for i*2 < n {
		c1 := i*2 + 1
		c2 := i*2 + 2
		max := c1

		if (c2 < n) && (s.Less(c1, c2)) {
			max = c2
		}
		if s.Less(max, i) {
			return
		}
		s.Swap(i, max)
		i = max
	}
}
