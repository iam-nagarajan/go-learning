package main

import "fmt"

func main() {
	var nums [3]int
	fmt.Println("Nums = ", nums)
	for i,n := range nums {
		fmt.Println(i, n)
	}
	withnums := [2]int{1,2}
	for _,n := range withnums {
		fmt.Println(n)
	}
	withnosize := [...]string{"A", "B", "C"}
	fmt.Println(len(withnosize))
	fmt.Printf("Capacity = %d\n", cap(withnosize))
	fmt.Println(withnosize[:3]) // Index cannot be negative
	// fmt.Println(withnosize[3:1]) Invalid
	// fmt.Println(withnosize[:67]) Out of bounds
	another := []int{1,2,3}
	fmt.Println(another)
	//fmt.Printf("%t, %t\n", nums, another)
	var noelement = make([]int, 5)
	//noelement[0] = 1
	fmt.Println(noelement)
	//nums = append(nums, 43) Not possible, cannot append to array
	numslice := append(nums[:], 43)
	fmt.Println("After append = ", numslice)
	fmt.Println("Setting nums[0] as 55")
	nums[0] = 55
	fmt.Println("Nums = ", nums)
	fmt.Println("NUmslice = ", numslice)
	fmt.Println("Setting numslice[1] as 66")
	numslice[1] = 66
	fmt.Println("Nums = ", nums)
	fmt.Println("Numslice = ", numslice)
	// Reslicing numslice
	reslice := make([]int, 0)
	// reslice = append(reslice, numslice[1:4]) // Throws error, add ... in the end to avoid error
	reslice = append(reslice, numslice[1:4]...) // Equivalent to copy
	//copy(reslice, numslice[1:4]) // Will have a copy
	fmt.Printf("Len = %d, Cap = %d\n", len(reslice), cap(reslice))
	reslice[1] = 77
	fmt.Println("numslice = ", numslice)
	fmt.Println("reslice = ", reslice)
}
