package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{29,19,28} 
	fmt.Println("---- arrays ----")

	var ages = [3]int{29,19,28} 


	names := [4]string{"yoshi","mario","peach","bowser"}
	names[2] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	fmt.Println("---- slices ----")
	
	var scores = []int{100,80,60}
	scores[2] = 99
	scores = append(scores,85)

	fmt.Println(scores,len(scores))

	// slice ranges 
	fmt.Println("---- slice ranges ----")

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo,rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}