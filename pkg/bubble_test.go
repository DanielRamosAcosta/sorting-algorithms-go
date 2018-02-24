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

func BenchmarkBubbleSortRandom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewRandom(DEFAULT_SIZE, int64(n+DEFAULT_SEED))
		aedasort.BubbleSort(s)
	}
}

func BenchmarkBubbleSortNearlySorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewNearlySorted(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_SPARSE)
		aedasort.BubbleSort(s)
	}
}

func BenchmarkBubbleSortReversed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewReversed(DEFAULT_SIZE)
		aedasort.BubbleSort(s)
	}
}

func BenchmarkBubbleSortFewUnique(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewFewUnique(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_GROUP)
		aedasort.BubbleSort(s)
	}
}
