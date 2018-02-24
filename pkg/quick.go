package aedasort

import "sort"

// QuickSort sorts the given sortable item using QuickSort Sort algorithm
// https://en.wikipedia.org/wiki/Quicksort
func QuickSort(s sort.Interface) {
	quickSort(s, 0, s.Len()-1)
}

func quickSort(s sort.Interface, beg int, end int) {
	b := beg
	e := end
	pi := (b + e) / 2

	for b < e {
		for s.Less(b, pi) {
			b++
		}
		for s.Less(pi, e) {
			e--
		}
		if b <= e {
			// Si cambiamos donde estaba el pivote, hay que actualizar su posición
			if b == pi {
				pi = e
			} else if e == pi {
				pi = b
			}
			s.Swap(b, e)
			b++
			e--
		}
	}
	if beg < e {
		quickSort(s, beg, e)
	}
	if b < end {
		quickSort(s, b, end)
	}
}
