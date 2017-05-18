package main

import "fmt"

func main() {
	fmt.Println("Usual loop")
	for i:=0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Multi var loop")
	for i, j := 0, 10; i < 10 && j > 0; i, j = i + 1, j-1 {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}

	fmt.Println("Loop with boolean")
	canRun := true
	count := 0
	for canRun {
		count++
		if count > 3 {
			canRun = false
		}
		fmt.Println("Can run is true")
	}
}
