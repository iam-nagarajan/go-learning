package main

import "fmt"

func main() {
	idToName := make(map[int]string)
	idToName[100486] = "Owler"
	idToName[100651] = "Google"
	idToName[100978] = "Microsoft"
	// idToName = {100980:"Oracle"} // Not possible
	anotherMap := map[int]string{100980:"Oracle"}
	fmt.Println(idToName)
	fmt.Println(anotherMap)
	cName := idToName[100651]
	fmt.Println(cName)
	name, isExist := idToName[100651]
	fmt.Println(name)
	fmt.Println(isExist)
	aName := idToName[100978]
	fmt.Println(aName)
}
