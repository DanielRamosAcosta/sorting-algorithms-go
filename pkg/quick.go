package aedasort

// QuickSort sorts the given sortable item using QuickSort Sort algorithm
// https://en.wikipedia.org/wiki/Quicksort
func QuickSort(s Interface) {
	quickSort(s, 0, s.Len()-1)
}

func quickSort(s Interface, beg int, end int) {
	b := beg
	e := end
	p := s.GetKey((b + e) / 2)

	for b < e {
		for s.GetKey(b) < p {
			b++
		}
		for s.GetKey(e) > p {
			e--
		}
		if b <= e {
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
