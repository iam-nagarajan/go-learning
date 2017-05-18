package main

import "fmt"

const con = 15

func main() {
	var str1 string = "Test"
	var str2 = "1"
	var int1 int
	var str string
	var f = 7.0
	fmt.Printf("str1 = %T, str2 = %T, int1 val = %d, str val = %s, f = %T\n", str1, str2, int1, str, f)
	fmt.Println(con)
	fmt.Printf("%T", con)
	//con = 16 // Cannot assign to const
	//fmt.Println(con)
}
