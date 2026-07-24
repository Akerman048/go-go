package main

import "fmt"


func main() {
myBill := newBill("Mario's bill")

myBill.addItem("onion soup", 4.60)
myBill.addItem("meat pie", 9.85)
myBill.addItem("toffee pudding", 4.95)
myBill.addItem("coffee", 3.25)
myBill.updateTip(10)

fmt.Println(myBill.format())
}
