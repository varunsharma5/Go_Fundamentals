package main

import "fmt"

func main() {

	age := 45

	fmt.Println(age < 45)

	if age < 45 {
		fmt.Println("Age is less than 45")
	} else if age > 45 {
		fmt.Println("Age is greater than 45")
	} else {
		fmt.Println("Age is 45")
	}

	names := []string{"name0", "name1", "name2", "name3", "name4"}

	for index, value := range names {
		if index == 2 {
			fmt.Println("Skipping position 1")
			continue
		}

		fmt.Printf("the value at position %v is %v\n", index, value)
	}
}
