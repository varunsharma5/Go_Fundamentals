package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 3.99, "cake": 5.99},
		tip:   0,
	}

	return b
}

// format the bill: Reciever function, binding this function with 'bill'
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v\t...$%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%v\t...$%0.2f", "total:", total)
	return fs
}
