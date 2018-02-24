package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestInsertionSort(t *testing.T) {
	s := SortType{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	aedasort.InsertionSort(s)
	failIfNotSorted(s, t)
}
