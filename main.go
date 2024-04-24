package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(propmt string, r *bufio.Reader) (string, error) {
	fmt.Print(propmt + ": ")
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name", reader)

	myBill := newBill(name)
	fmt.Println(">>> Created the bill - ", myBill.name)

	myBill = newBill("varun's bill")
	return myBill
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip, q - quit)", reader)

	switch opt {
	case "a":
		name, _ := getInput("Enter Item Name", reader)
		price, _ := getInput("Enter item price", reader)
		price_float, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("value must be an integer")
			promptOptions(b)
		}
		b.addItem(name, price_float)
		fmt.Println(">>> Item Added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip ($)", reader)
		tip_float, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("value must be an integer")
			promptOptions(b)
		}
		b.updateTip(tip_float)
		fmt.Println(">>> Tip Added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
	case "q":
		fmt.Println("Bye Bye")
		return
	default:
		fmt.Println("invalid option...!")
		promptOptions(b)
	}
}

func main() {
	var myBill bill = createBill()
	promptOptions(myBill)
	fmt.Println(myBill.format())

}
