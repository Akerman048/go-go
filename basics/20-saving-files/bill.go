package main

import (
	"fmt"
	"os"
	"path/filepath"
)


type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make new bills
func newBill(name string) bill {
b := bill{
	name: name,
	items: map[string]float64{},
	tip: 0,
}

return b
}

// format the bill

func (b bill) format() string {
fs := "Bill breakdown: \n"
var total float64 = 0

// list items
for k, v := range b.items {
	fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
	total += v
	
}

// add tip
fs += fmt.Sprintf("%-25v ...$%v\n","tip:", b.tip)

// total 
fs += fmt.Sprintf("%-25v ...$%0.2f","total:", total+b.tip)

return  fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64){
	b.items[name] = price
}


// save bill
func (b *bill) save() error {
err := os.MkdirAll("bills", 0755)

if err != nil {
	return err
}

filename := b.name + ".txt"
filepath := filepath.Join("bills", filename)

data := []byte(b.format())

err = os.WriteFile(filepath, data, 0644)
if err != nil{
	return err
}

fmt.Println("Bill was saved to:", filepath)
return nil
}