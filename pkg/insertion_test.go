package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestInsertionSort(t *testing.T) {
	s := NewRandom(DEFAULT_SIZE, DEFAULT_SEED)
	aedasort.InsertionSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkInsertionSortRandom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewRandom(DEFAULT_SIZE, int64(n+DEFAULT_SEED))
		aedasort.InsertionSort(s)
	}
}

func BenchmarkInsertionSortNearlySorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewNearlySorted(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_SPARSE)
		aedasort.InsertionSort(s)
	}
}

func BenchmarkInsertionSortReversed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewReversed(DEFAULT_SIZE)
		aedasort.InsertionSort(s)
	}
}

func BenchmarkInsertionSortFewUnique(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewFewUnique(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_GROUP)
		aedasort.InsertionSort(s)
	}
}
