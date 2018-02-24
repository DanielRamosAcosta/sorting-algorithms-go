package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestSelectionSort(t *testing.T) {
	s := NewRandom(DEFAULT_SIZE, DEFAULT_SEED)
	aedasort.SelectionSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkSelectionSortRandom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewRandom(DEFAULT_SIZE, int64(n+DEFAULT_SEED))
		aedasort.SelectionSort(s)
	}
}

func BenchmarkSelectionSortNearlySorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewNearlySorted(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_SPARSE)
		aedasort.SelectionSort(s)
	}
}

func BenchmarkSelectionSortReversed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewReversed(DEFAULT_SIZE)
		aedasort.SelectionSort(s)
	}
}

func BenchmarkSelectionSortFewUnique(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewFewUnique(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_GROUP)
		aedasort.SelectionSort(s)
	}
}
