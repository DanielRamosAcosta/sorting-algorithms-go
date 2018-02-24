package aedasort_test

import (
	"sort"
	"testing"
)

func BenchmarkNativeSort(b *testing.B) {
	BenchAlgorithm(b, sort.Sort)
}
