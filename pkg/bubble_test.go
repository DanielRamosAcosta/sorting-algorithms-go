package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestBubbleSort(t *testing.T) {
	s := NewRandom(DEFAULT_SMALL_SIZE, DEFAULT_SEED)
	aedasort.BubbleSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkBubbleSort(b *testing.B) {
	BenchAlgorithm(b, aedasort.BubbleSort)
}
