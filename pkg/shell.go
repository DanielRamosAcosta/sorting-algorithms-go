package aedasort

import (
	"sort"
)

const shellSortAlpha = 0.5

// ShellSort sorts the given sortable item using Shell Sort algorithm
// https://en.wikipedia.org/wiki/Shellsort
func ShellSort(s sort.Interface) {

	delta := s.Len()

	for delta > 1 {
		delta = int(float32(delta) * shellSortAlpha)

		for i := delta; i < s.Len(); i++ {
			it := i
			var j int

			for j = i; (j >= delta) && (s.Less(it, j-delta)); j -= delta {
				if j == it {
					it = j - delta
				} else if (j - delta) == it {
					it = j
				}
				s.Swap(j, j-delta)
			}

			s.Swap(j, it)
		}
	}
}
