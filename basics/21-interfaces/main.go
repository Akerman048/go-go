package main

import (
	"fmt"
	"math"
)


type saver interface {
	save()
}

type bill struct {
	name string
}

func (b *bill) save(){
	fmt.Println("Saving bill:", b.name)
}

func saveData(s saver) {
	s.save()
}

//----------------


// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return  s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}


// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return  2 * math.Pi * c.radius
}

func main() {
myBill := bill{name: "Leon"}

saveData(&myBill)
}
