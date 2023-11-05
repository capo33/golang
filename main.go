package main

import "fmt"

func main() {
	// // Strings
	// var nameOne string = "John"
	// var nameTwo = "Sally"
	// var nameThree string

 	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "Michael"
	// nameThree = "Jennifer"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// // Integers
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40 

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // Bits & Memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint = 25

	// fmt.Println(numOne, numTwo, numThree)

	// // Floats
	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 356789.123456789
	// scoreThree := 1.5

	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	age := 25
	name := "Capo"

	// Print Type
	fmt.Print("Hello, \n")
	fmt.Print("Capo! \n")

	// Print Line
	fmt.Println("Hello, Capo!")
	fmt.Println("Goodbye, Capo!")
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings)
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age,name)
	fmt.Printf("age is of type %T \n", name)
	fmt.Printf("You scored %f points! \n", 225.55)
	fmt.Printf("You scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)	
}
