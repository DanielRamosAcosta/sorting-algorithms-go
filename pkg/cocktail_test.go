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

func BenchmarkCocktailSort(b *testing.B) {
	BenchAlgorithm(b, aedasort.CocktailSort)
}
