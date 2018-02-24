package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestInsertionSort(t *testing.T) {
	s := NewRandom(DEFAULT_SMALL_SIZE, DEFAULT_SEED)
	aedasort.InsertionSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkInsertionSort(b *testing.B) {
	BenchAlgorithm(b, aedasort.InsertionSort)
}
