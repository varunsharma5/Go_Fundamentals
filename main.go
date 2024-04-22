package main

import "fmt"

func main() {
	//           key    value
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 3.55,
	}

	fmt.Println(menu)

	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Printf("Key: %v, value: %0.2f\n", k, v)
	}

	// int as key type
	phonebook := map[int]string{
		234: "Book1",
		235: "Book2",
		236: "Book3",
	}

	fmt.Println(phonebook)

	phonebook[234] = "Book_new"
	fmt.Println(phonebook)

	// map of map
	map1 := map[int]map[int]string{
		126328: phonebook,
		// 343432: menu,
	}

	fmt.Println(map1)

}
