package go_sorts

import (
	//"fmt"
	//"internal/testenv"
	"math"
	"math/rand"
	//. "sort"
	"sort"
	//"strconv"
	//stringspkg "strings"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestSortIntSlice(t *testing.T) {
	data := ints
	sorted := sort.IntSlice(data[0:])
	sort.Sort(sorted)

	selectionData := ints
	selection := sort.IntSlice(selectionData[0:])
	SelectionSort(selection)

	if !sort.IsSorted(selection) {
		t.Errorf("original %v", ints)
		t.Errorf("sorted %v", sorted)
		t.Errorf("   got %v", selection)
	}


}

func TestSortFloat64Slice(t *testing.T) {
	data := float64s
	a := sort.Float64Slice(data[0:])
	sort.Sort(a)

	selectionData := float64s
	selection := sort.Float64Slice(selectionData[0:])
	SelectionSort(selection)

	if !sort.IsSorted(a) {
		t.Errorf("original %v", float64s)
		t.Errorf("sorted %v", data)
		t.Errorf("   got %v", selection)
	}
}

func TestSortStringSlice(t *testing.T) {
	data := strings
	a := sort.StringSlice(data[0:])
	sort.Sort(a)

	selectionData := strings
	selection := sort.StringSlice(selectionData[0:])
	SelectionSort(selection)

	if !sort.IsSorted(a) {
		t.Errorf("original %v", strings)
		t.Errorf("sorted %v", data)
		t.Errorf("   got %v", selection)
	}
}

////not implement yet
//func TestInts(t *testing.T) {
//	data := ints
//	sort.Ints(data[0:])
//
//	if !sort.IntsAreSorted(data[0:]) {
//		t.Errorf("sorted %v", ints)
//		t.Errorf("   got %v", data)
//	}
//}
//
////not implement yet
//func TestFloat64s(t *testing.T) {
//	data := float64s
//	sort.Float64s(data[0:])
//	if !sort.Float64sAreSorted(data[0:]) {
//		t.Errorf("sorted %v", float64s)
//		t.Errorf("   got %v", data)
//	}
//}
//
////not implement yet
//func TestStrings(t *testing.T) {
//	data := strings
//	sort.Strings(data[0:])
//	if !sort.StringsAreSorted(data[0:]) {
//		t.Errorf("sorted %v", strings)
//		t.Errorf("   got %v", data)
//	}
//}

//func TestSlice(t *testing.T) {
//	data := strings
//	sort.Slice(data[:], func(i, j int) bool {
//		return data[i] < data[j]
//	})
//	if !sort.SliceIsSorted(data[:], func(i, j int) bool { return data[i] < data[j] }) {
//		t.Errorf("sorted %v", strings)
//		t.Errorf("   got %v", data)
//	}
//}

func TestSortLarge_Random(t *testing.T) {
	n := 1000000
	if testing.Short() {
		n /= 100
	}
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}
	if sort.IntsAreSorted(data) {
		t.Fatalf("terrible rand.rand")
	}
	sort.Ints(data)

	selection := sort.IntSlice(data[0:])
	SelectionSort(selection)

	if !sort.IsSorted(selection) {
		//t.Errorf("original %v", ints)
		//t.Errorf("sorted %v", sorted)
		//t.Errorf("   got %v", selection)
		//t.Log('Selection Sort ')
		t.Errorf("Selection Sort: sort didn't sort - 1M ints")
	}

	if !sort.IntsAreSorted(data) {
		t.Errorf("sort didn't sort - 1M ints")
	}
}