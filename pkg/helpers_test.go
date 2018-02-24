package aedasort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/danielramosacosta/sorting-algorithms-go/pkg"
)

type SortType []uint

const DEFAULT_SIZE = 10000
const DEFAULT_SMALL_SIZE = 15
const DEFAULT_SEED = 7
const DEFAULT_SPARSE = 10
const DEFAULT_GROUP = 50

func (a SortType) Len() int              { return len(a) }
func (a SortType) Swap(i, j int)         { a[i], a[j] = a[j], a[i] }
func (a SortType) Less(i, j int) bool    { return a[i] < a[j] }
func (a SortType) Greater(i, j int) bool { return a[i] > a[j] }
func (a SortType) GetKey(i int) uint     { return a[i] }

func New(length uint) SortType {
	s := make(SortType, length)

	for i := range s {
		s[i] = uint(i) + 1
	}

	return s
}

func NewRandom(length uint, seed int64) SortType {
	s := rand.NewSource(seed)
	r := rand.New(s)
	st := New(length)

	r.Shuffle(len(st), func(i, j int) {
		st[i], st[j] = st[j], st[i]
	})

	return st
}

func NewReversed(length uint) SortType {
	s := make(SortType, length)

	for i := range s {
		s[i] = length - uint(i)
	}

	return s
}

func NewNearlySorted(length uint, seed int64, sparse uint) SortType {
	s := rand.NewSource(seed)
	r := rand.New(s)
	st := New(length)

	for i := 0; uint(i) < length; i++ {
		newIndex := uint(i + r.Intn(int(sparse)))

		if newIndex < length {
			st[i], st[newIndex] = st[newIndex], st[i]
		}
	}

	return st
}

func NewFewUnique(length uint, seed int64, groups uint) SortType {
	s := rand.NewSource(seed)
	r := rand.New(s)
	st := make(SortType, length)

	for i := uint(0); i < length; i += groups {
		for j := i; j < length; j++ {
			st[j] = i + 1
		}
	}

	r.Shuffle(len(st), func(i, j int) {
		st[i], st[j] = st[j], st[i]
	})

	return st
}

func NewOfType(t string, length uint, seed int64) SortType {
	switch t {
	case "random":
		return NewRandom(length, seed)
	case "nearlysorted":
		return NewNearlySorted(length, seed, length/10)
	case "reversed":
		return NewReversed(length)
	case "fewunique":
		return NewFewUnique(length, seed, (length/10)+1)
	}
	panic(fmt.Sprintf("Unknown type %v", t))
}

func failIfNotSorted(a SortType, t *testing.T) {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			t.Error("Is not sorted", a)
		}
	}
}

func BenchAlgorithm(b *testing.B, f func(sort.Interface)) {
	benchTypes := []string{
		"random",
		"nearlysorted",
		"reversed",
		"fewunique",
	}

	slicesLengths := []uint{10, 50, 100, 200, 300, 600, 1000, 2000, 5000}

	for _, t := range benchTypes {
		b.Run(t, func(b *testing.B) {
			for _, n := range slicesLengths {
				b.Run(fmt.Sprintf("%v", n), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						slice := NewOfType(t, n, int64(int(n)*b.N))
						f(slice)
					}
				})
			}
		})
	}
}

func BenchAlgorithmCustomInterface(b *testing.B, f func(aedasort.Interface)) {
	benchTypes := []string{
		"random",
		"nearlysorted",
		"reversed",
		"fewunique",
	}

	slicesLengths := []uint{10, 50, 100, 200, 300, 600, 1000, 2000, 5000}

	for _, t := range benchTypes {
		b.Run(t, func(b *testing.B) {
			for _, n := range slicesLengths {
				b.Run(fmt.Sprintf("%v", n), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						slice := NewOfType(t, n, int64(int(n)*b.N))
						f(slice)
					}
				})
			}
		})
	}
}
