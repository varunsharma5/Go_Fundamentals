package main

import "fmt"

func main() {
	var myBill bill
	myBill = newBill("varun's bill")

	myBill.addItem("cake", 5.99)
	myBill.addItem("pie", 3.99)
	myBill.updateTip(10)
	fmt.Println(myBill.format())

	// myBill.format()

}
