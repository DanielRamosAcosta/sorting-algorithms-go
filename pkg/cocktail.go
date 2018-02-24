package aedasort

import (
	"sort"
)

// CocktailSort sorts the given sortable item using Cocktail Sort algorithm
// https://en.wikipedia.org/wiki/Cocktail_shaker_sort
func CocktailSort(s sort.Interface) {
	change := true
	for beg, end := 1, s.Len()-1; (beg < end) && change; {
		for j := beg; j <= end; j++ {
			if s.Less(j, j-1) {
				s.Swap(j-1, j)
				change = true
			}
		}
		end--

		for j := end; j >= beg; j-- {
			if s.Less(j, j-1) {
				s.Swap(j-1, j)
				change = true
			}
		}

		beg++
	}
}
