package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestInsertionSort(t *testing.T) {
	s := NewRandom(10, 7)
	aedasort.InsertionSort(s)
	failIfNotSorted(s, t)
}
