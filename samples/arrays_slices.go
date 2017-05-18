package main

import "fmt"

func main() {
	var nums [3]int
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
}
