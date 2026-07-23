package main

import "fmt"


func main() {

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	fmt.Println("---- looping maps ----")

	for k,v := range menu{
		fmt.Println(k, "-", v)
	}

	// ints as key type
	fmt.Println("---- ints as key type ----")

	phonebook := map[int]string{
		23569345:"Koni",
		92346235:"Armin",
		63731456:"Mikasa",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[63731456])

	phonebook[92346235] = "Jacques"
	fmt.Println(phonebook)

	phonebook[23569345] = "Hanzhi"
	fmt.Println(phonebook)
}
