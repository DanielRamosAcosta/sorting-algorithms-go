package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestQuickSort(t *testing.T) {
	s := NewRandom(DEFAULT_SMALL_SIZE, DEFAULT_SEED)
	aedasort.QuickSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkQuickSort(b *testing.B) {
	BenchAlgorithmCustomInterface(b, aedasort.QuickSort)
}
