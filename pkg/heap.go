package aedasort

import (
	"sort"
)

// HeapSort sorts the given sortable item using Heap Sort algorithm
// https://en.wikipedia.org/wiki/Heapsort
func HeapSort(s sort.Interface) {
	for i := (s.Len() - 1) / 2; i >= 0; i-- {
		siftDown(s, i, s.Len())
	}

	for i := s.Len() - 1; i >= 0; i-- {
		s.Swap(0, i)
		siftDown(s, 0, i)
	}
}

func siftDown(s sort.Interface, i, n int) {
	root := i
	for {
		child := 2*root + 1
		if child >= n {
			break
		}
		if (child+1 < n) && s.Less(child, child+1) {
			child++
		}
		if !s.Less(root, child) {
			return
		}
		s.Swap(root, child)
		root = child
	}
}
