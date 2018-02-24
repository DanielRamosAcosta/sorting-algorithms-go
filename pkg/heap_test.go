package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestHeapSort(t *testing.T) {
	s := NewRandom(DEFAULT_SMALL_SIZE, DEFAULT_SEED)
	aedasort.HeapSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkHeapSort(b *testing.B) {
	BenchAlgorithm(b, aedasort.HeapSort)
}
