package aedasort

import (
	"sort"
)

type Interface interface {
	sort.Interface
	GetKey(int) uint
}
