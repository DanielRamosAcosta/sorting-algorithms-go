package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestSelectionSort(t *testing.T) {
	s := NewRandom(DEFAULT_SMALL_SIZE, DEFAULT_SEED)
	aedasort.SelectionSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkSelectionSort(b *testing.B) {
	BenchAlgorithm(b, aedasort.SelectionSort)
}
