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

