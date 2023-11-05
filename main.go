package main

import "fmt"

func main() {
	// Arrays, always a fixed size and type
	var ages [3]int = [3]int{20, 25, 30}
	var num = [5]int{1, 2, 3, 4, 5}

	name := [4]string{"John", "Paul", "George", "Ringo"}

	// change value of name array
	name[1] = "Bob"

	fmt.Println(ages, len(ages))
	fmt.Println(num, len(num))
	fmt.Println(name, len(name))

	// Slices	- use arrays under the hood
	var scores = []int{100, 50, 60}
	scores[2] = 25
	
	fmt.Println(scores, len(scores))

	var appendedSlice = append(scores, 85)
	fmt.Println(appendedSlice, len(appendedSlice))

	fmt.Println(scores, len(scores))

	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := name[1:3]
	rangeTwo := name[2:]
	rangeThree := name[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
