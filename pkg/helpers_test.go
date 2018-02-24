package aedasort_test

import "testing"

type SortType []int

func (a SortType) Len() int           { return len(a) }
func (a SortType) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortType) Less(i, j int) bool { return a[i] < a[j] }
func (a SortType) GetDebug(i int) int { return a[i] }

func failIfNotSorted(a SortType, t *testing.T) {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			t.Error("Is not sorted", a)
		}
	}
}
