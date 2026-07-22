package main

import "fmt"


func main() {
	age := 29
	name := "Valerii"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")


	// Println
	fmt.Println("hello Akerman!")
	fmt.Println("see you Akerman!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings)

	// %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age,name)
	// %q = quote
	fmt.Printf("my age is %v and my name is %q", age,name)
	// %T = type of var
	fmt.Printf("age is of type %T", age)
	// %f = float32 or float64.
	fmt.Printf("you scored %f points! \n", 28.88)
	// %.1f = how many digits after .
	fmt.Printf("you scored %.1f points! \n", 28.88)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n",age,name)
	fmt.Println("the saved string is:",str)
}