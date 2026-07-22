package main

import "fmt"

func main() {
	const appName = "Simple Calculator"

	var firstNumber float64
	var secondNumber float64

	fmt.Println(appName)
	fmt.Println("-----------------")

	fmt.Print("Enter first number: ")
	_, firstErr := fmt.Scan(&firstNumber)

	fmt.Print("Enter second number: ")
	_, secondErr := fmt.Scan(&secondNumber)

	fmt.Println()
	fmt.Println("Input errors:")
	fmt.Println("First input error:", firstErr)
	fmt.Println("Second input error:", secondErr)

	fmt.Println()
	fmt.Printf("%.2f + %.2f = %.2f\n",
		firstNumber,
		secondNumber,
		firstNumber+secondNumber,
	)

	fmt.Printf("%.2f - %.2f = %.2f\n",
		firstNumber,
		secondNumber,
		firstNumber-secondNumber,
	)

	fmt.Printf("%.2f * %.2f = %.2f\n",
		firstNumber,
		secondNumber,
		firstNumber*secondNumber,
	)

	fmt.Printf("%.2f / %.2f = %.2f\n",
		firstNumber,
		secondNumber,
		firstNumber/secondNumber,
	)
}