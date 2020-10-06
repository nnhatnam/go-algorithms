package go_sorts

import (
	"sort"
)
//private function won't handle edge cases
func selectionSort(data sort.Interface, a , b int){

	if b - a > 1 {
		for i := a ; i < b; i++ {
			min := i
			for j := i+1; j < b ; j++{
				if data.Less(j , min) { //if j should before min, min = j
					min = j
				}
			}
			if min != i {
				data.Swap(min, i)
			}

		}
	}
}

func SelectionSort(data sort.Interface) {
	selectionSort(data, 0, data.Len())
}

func insertionSort(data sort.Interface, a , b int) {

	if b - a > 1 {
		for i := a + 1; i < b; i++{
			for j := i; j > a && data.Less(j,i); j-- {
				data.Swap(j , j-1 )
			}
		}
	}

}

func InsertionSort(data sort.Interface) {
	insertionSort(data, 0, data.Len())
}

