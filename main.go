package main

import "fmt"

func main() {
	var myBill bill
	myBill = newBill("varun's bill")

	fmt.Println(myBill)
	fmt.Printf("%+v\n", myBill)

}
