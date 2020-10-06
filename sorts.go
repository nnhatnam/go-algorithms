package go_sorts

import (
	"sort"
)

func selectionSort(data sort.Interface, a , b int){

	if b - a > 1 {
		//fmt.Println(data)
		for i:=a ; i < b; i++ {
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