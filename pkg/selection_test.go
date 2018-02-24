package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestSelectionSort(t *testing.T) {
	s := NewRandom(100, 7)
	aedasort.SelectionSort(s)
	failIfNotSorted(s, t)
}
