package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("7.0 / 3 = ", 7.0/3)
	fmt.Println("10/3 = ", 10/3)
	//fmt.Println("String + Boolean = ", "str" + true) //Error
	fmt.Println("Convert primitive to str = ", strconv.FormatBool(true))
}
