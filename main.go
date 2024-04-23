package main

import "fmt"

func updateMenu(menu map[string]float64) {
	menu["orange"] = 6.00
}

func main() {

	// Group 2: map, slice, function

	menu := map[string]float64{
		"pie":   4.99,
		"apple": 3.00,
	}

	updateMenu(menu)

	fmt.Println(menu)

}
