package aedasort_test

import (
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

func TestShellSort(t *testing.T) {
	s := NewRandom(DEFAULT_SMALL_SIZE, DEFAULT_SEED)
	aedasort.ShellSort(s)
	failIfNotSorted(s, t)
}

func BenchmarkShellSort(b *testing.B) {
	BenchAlgorithm(b, aedasort.ShellSort)
}
