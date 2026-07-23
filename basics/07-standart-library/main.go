package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("---- strings ----")
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "arigato"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting,"ll"))
	fmt.Println(strings.Split(greeting,"e"))

	// the original string value
	fmt.Println("original string value =", greeting)

	fmt.Println("---- sort ----")
	ages := []int{45,22,36,28,88,71,53}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages,36)
	fmt.Println(index)

	names := []string{"yoshi","mario","peach","bowser","luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names,"bowser"))
}