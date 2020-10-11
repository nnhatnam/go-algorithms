package sorts

import "sort"

type Interface interface {
	sort.Interface
	Elem(i int) interface{}
}

