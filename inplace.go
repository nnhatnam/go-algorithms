package sorts

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



//func shellSortWGap(data sort.Interface, a , b, gap int){
//	if b - a > 1 {
//		if a + gap < b {
//			for i := a + gap; i < b; i++ {
//				if data.Less(i , i - gap) {
//					data.Swap(i, i - gap )
//				}
//			}
//		}
//		insertionSort(data, a , b)
//	}
//}

//h-sort step of shell sort
//with insertion sort, h = 1
//call this function only if it's guaranteed that a - b > 1
func hsort(data sort.Interface, a , b, gap int) {
	for i := a + gap; i < b; i++ {
		for j := i; j >= gap && data.Less(j, j - gap); j-= gap {
			data.Swap(j, j - gap)
		}
	}
}

//invariant: binary insertion, using a sentinel
func insertionSort(data sort.Interface, a , b int) {

	if b - a > 1 {
		hsort(data, a, b, 1)
	}

}

func InsertionSort(data sort.Interface) {
	insertionSort(data, 0, data.Len())
}

// Knuth Sequence
func shellSort(data sort.Interface, a , b int){
	if b - a > 1 {

		// find Knuth Sequence
		h := 1
		for ; h < (b - a) / 3 ; {
			h = 3*h + 1
		}

		for ; h >= 1 ; h = (h - 1) / 3 {
			for i := a + h; i < b; i++ {
				for j := i; j >= h && data.Less(j, j - h); j -= h {
					data.Swap(j, j - h)
				}
			}
		}

	}
}

func ShellSort(data sort.Interface) {
	shellSort(data, 0, data.Len())
}

//implement a max heapify
// first------------lo------hi------------------last
func siftDown(data sort.Interface, lo , hi , first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child + 1 < hi && data.Less(first + root, first + child + 1) {
			data.Swap(first + root, first + child)
		}

		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first + root, first + child)
		root = child

		//if data.Less(root, lChild ) {
		//	data.Swap(root, lChild)
		//}
	}
}

// b is length from 0 to index b - 1
func heapSort(data sort.Interface, a , b int){
	//parent index = i => left = 2*i + 1 , right = 2*i + 2
	//child index = i => parent = (i - 1) / 2

	//build a heap
	//heapsize = b
	first := a				// root index
	lo := 0					// root offset
	hi := b - a	            // last leaf offset

	//build a heap
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	for i := hi - 1; i >= 0 ; i-- {
		data.Swap(first, first + i)
		siftDown(data, lo, hi - i , first)
	}
}

func HeapSort(data sort.Interface) {
	heapSort(data, 0, data.Len())
}

// b is length from 0 to index b - 1
//func mergeSort(data sort.Interface, a, b int) {
//	lo := 0
//	n := b / 2
//
//	mergeSort(data, )
//}


func GnomeSort(data sort.Interface) {

		//i :  item after the in order items
		//j :  pointer to next value
		for i, j := 1 ,2  ; i < data.Len(); {
			if data.Less(i, i - 1) {
				data.Swap(i , i - 1)
				i--
				if i > 0 {
					continue
				}
			}

			i = j
			j = i + 1
		}
}

func BubbleSort(data sort.Interface) {

	for i := 0; i < data.Len(); i++{
		swapped  := false
		for j:= 1; i < data.Len(); j++{
			if data.Less(j, j - 1) {
				data.Swap(j, j - 1)
				swapped  = true
			}
		}
		if !swapped  {
			return
		}
	}
}

func CombSort(data sort.Interface) {
	swapped := false
	for gap := data.Len(); gap > 1 || swapped ; {

		if gap > 1 {
			//gap = gap / 1.3
			gap = gap * 10 / 13
		}

		for i := gap; i < data.Len(); i++ {
			if data.Less(i - gap, i) {
				data.Swap(i - gap, i )
				swapped = true
			}
		}
	}
}

func OddEvenSort(data sort.Interface) {
	isSort := false
	for ; !isSort ; {
		isSort = true
		for i := 2 ; i < data.Len() - 1; i += 2 {
			if data.Less(i, i - 2 ) {
				data.Swap(i , i - 2)
				isSort = false
			}
		}

		for i := 3 ; i < data.Len() - 1; i += 2 {
			if data.Less(i, i - 2 ) {
				data.Swap(i , i - 2)
				isSort = false
			}
		}

	}
}

func CocktailSort(data sort.Interface) {
	start := 1
	//do - while
	//should not use for loop this way
	for  flag := false  ; !flag ; flag , start = !flag , start + 1  {
		for i := start ; i < data.Len() ; i++ {
			if data.Less(i , i-1) {
				data.Swap(i, i - 1)
				flag = true
			}
		}

		if !flag {
			return
		}

		for i := data.Len() - 1 - start; i > start ; i-- {
			if data.Less(i, i - 1 ) {
				data.Swap(i , i - 1)
				flag = true
			}
		}

	}
}