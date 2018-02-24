package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestCocktailSort(t *testing.T) {
	s := NewRandom(DEFAULT_SMALL_SIZE, DEFAULT_SEED)
	aedasort.CocktailSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkCocktailSortRandom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewRandom(DEFAULT_SIZE, int64(n+DEFAULT_SEED))
		aedasort.CocktailSort(s)
	}
}

func BenchmarkCocktailSortNearlySorted(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewNearlySorted(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_SPARSE)
		aedasort.CocktailSort(s)
	}
}

func BenchmarkCocktailSortReversed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewReversed(DEFAULT_SIZE)
		aedasort.CocktailSort(s)
	}
}

func BenchmarkCocktailSortFewUnique(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := NewFewUnique(DEFAULT_SIZE, int64(n+DEFAULT_SEED), DEFAULT_GROUP)
		aedasort.CocktailSort(s)
	}
}
