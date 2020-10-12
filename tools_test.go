package sorts

import (
	"fmt"
	"sort"
	"testing"
)

var ints1 = [...]int{1 , 2 , 3 , 4, 5 , 6, 7 , 8 , 9, 10 , 11, 12, 13 }
var ints2 = [...]int{1 , 2 , 3 , 4, 5 , 6, 7 , 8 , 9 , 10 }

func swapRangeTest(data sort.Interface, a , b , n int) {
	//fmt.Println("inside swapRangeTest")
	//fmt.Println("before swap: ", data)
	for i := 0; i < n; i++ {
		//fmt.Println(fmt.Sprintf("swap %v and %v", a + i , b + i))
		data.Swap(a + i, b + i)
	}
	//fmt.Println("after swap: ", data)
}

// Rotate two consecutive blocks u = data[a:m] and v = data[m:b] in data:
// Data of the form 'x u v y' is changed to 'x v u y'.
// Rotate performs at most b-a many calls to data.Swap.
// Rotate assumes non-degenerate arguments: a < m && m < b.
func rotateTest(data sort.Interface, a, m, b int) {
	//m is index, not offset

	i := m - a // find the offset from a -> m
	j := b - m // find the offset from m -> b
	fmt.Println("data: ", data , a , m , b)
	fmt.Println("Step 1: i , j := " , i , j )
	for i != j {
		//fmt.Println("- inside i != j for loop, before swap: ", i , j , data)
		if i > j {
			fmt.Println("i , j : ", i , j )
			fmt.Println(fmt.Sprintf("-- swapRangeTest(%v , %v, %v , %v) ", "data", m - i , m , j ))
			fmt.Println("--- before swap: ",data)
			swapRangeTest(data, m-i, m, j)
			fmt.Println("--- after swap: ",data)
			i -= j
		} else {
			fmt.Println(fmt.Sprintf("-- swapRangeTest(%v , %v, %v , %v) ", "data", m - i , m+j-i , i ))
			fmt.Println("--- before swap: ",data)
			swapRangeTest(data, m-i, m+j-i, i)
			fmt.Println("--- after swap: ",data)
			j -= i
		}
		//fmt.Println("- inside i != j for loop, after swap: ", i , j , data)

	}
	// i == j
	fmt.Println("Step 2 : i == j so we swap with i, j = " , i , j )
	fmt.Println("--- before swap: ",data)
	swapRangeTest(data, m-i, m, i)
	fmt.Println("--- after swap: ",data)
}

func TestRotateBehavior(t *testing.T) {
	fmt.Println(ints1[0:3])
	data1 := sort.IntSlice(ints1[:])
	rotateTest(data1, 0 , 3 , data1.Len())


}
