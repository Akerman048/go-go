package main

import "fmt"

func main() {
	const app = "Go go"
	var age int = 29
	var name = "Valerii"
	var height float32 = 175.5
	isLearning := true
	newCountry, newCity := "Spain", "Barcelona"

	fmt.Println("App:", app)
	fmt.Println(newCountry, newCity)
	fmt.Println()

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.1f cm\n", height)
	fmt.Printf("Learning Go: %t\n", isLearning)

	fmt.Println()

	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of height: %T\n", height)
	fmt.Printf("Type of isLearning: %T\n", isLearning)

	age = 30

	fmt.Println()
	fmt.Printf("Age next year: %d\n", age)

	productName := "Go-go"
	price := 99.99
	quantity := 8
	isAvailable := true

	total := price * float64(quantity)

	fmt.Printf("Product: %s\n", productName)
	fmt.Printf("Price: %.0f\n", price)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Is available: %t\n", isAvailable)
	fmt.Printf("Total: %.1f\n", total)
}
