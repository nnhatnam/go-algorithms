package sorts

import "sort"

//swap data[a : a + n] and data[b:b+n]
//TODO: need more studied to find a better way
func swapRange(data sort.Interface, a , b , n int) {
	for i := 0; i < n; i++ {
		data.Swap(a + i, b + i)
	}
}

//Block-Swap algorithm
// Rotate two consecutive blocks u = data[a:m] and v = data[m:b] in data:
// Data of the form 'x u v y' is changed to 'x v u y'.
// Rotate performs at most b-a many calls to data.Swap.
// Rotate assumes non-degenerate arguments: a < m && m < b.
func rotate(data sort.Interface, a, m, b int) {
	//m is index, not offset

	i := m - a // find the offset from a -> m
	j := b - m // find the offset from m -> b

	for i != j {
		if i > j {
			swapRange(data, m-i, m, j)
			i -= j
		} else {
			swapRange(data, m-i, m+j-i, i)
			j -= i
		}
	}
	// i == j
	swapRange(data, m-i, m, i)
}

//https://github.com/BonzaiThePenguin/WikiSort/blob/master/Chapter%201.%20Tools.md
//reverse the items in range data[lo:hi]
func reversing(data sort.Interface, lo , hi int) {
	midOffset := (hi - lo) / 2
	for i := midOffset; i >= 0 ; i-- {
		data.Swap(lo + i, hi - i )
	}
}

//https://github.com/BonzaiThePenguin/WikiSort/blob/master/Chapter%201.%20Tools.md
//rotate using reversal algorithm 
func rotating(data sort.Interface, lo , hi , amount int) {
	reversing(data, lo , lo + amount - 1)
	reversing(data, lo + amount , hi)
	reversing(data, lo , hi)
}