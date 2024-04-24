package main

import (
	"fmt"
	"os"
)

// struct pointers are automatically dereferened
// in the below receiver funcation, upon passing struct pointer, we only need to
// chnage bill->*bill and no changes in required inside function as dereferencing is happing
// automatically by Go compiler

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill: Reciever function, binding this function with 'bill'
// (b *bill) - by adding this before function name, this functions is mapped (encapsulated) with the struc 'bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v\t...$%v\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v\t...$%0.2f\n", "tip:", b.tip)

	fs += fmt.Sprintf("%v\t...$%0.2f", "total:", total+b.tip)
	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
}
