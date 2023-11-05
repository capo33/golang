package main

import "fmt"

func main() {
	var nameOne string = "John"
	var nameTwo = "Sally"
	var nameThree string

 	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Michael"
	nameThree = "Jennifer"

	fmt.Println(nameOne, nameTwo, nameThree)
}
