package aedasort_test

import (
	"math/rand"
	"testing"
)

type SortType []uint

const DEFAULT_SIZE = 10000
const DEFAULT_SEED = 7
const DEFAULT_SPARSE = 10
const DEFAULT_GROUP = 50

func (a SortType) Len() int            { return len(a) }
func (a SortType) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a SortType) Less(i, j int) bool  { return a[i] < a[j] }
func (a SortType) GetDebug(i int) uint { return a[i] }

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

func failIfNotSorted(a SortType, t *testing.T) {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			t.Error("Is not sorted", a)
		}
	}
}
