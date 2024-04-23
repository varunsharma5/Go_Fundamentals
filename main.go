package main

import "fmt"

func updateName(m *string) {
	*m = "Name2"
}

func main() {

	name := "Name1"

	m := &name

	fmt.Println("memory location of the variable is: ", m)
	fmt.Println("value at memory address: ", *m)

	fmt.Println(name)

	updateName(m)

	fmt.Println(name)

}
