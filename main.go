package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v\n", name)
}

func cycleNames(names []string, fn func(string)) {
	for _, value := range names {
		fn(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * (r * r)
}

func main() {
	sayGreeting("Varun")
	cycleNames([]string{"name0", "name1", "name2"}, sayGreeting)

	area1 := circleArea(10.5)
	area2 := circleArea(15)

	fmt.Println(area1, area2)
	fmt.Printf("area1 is %0.2f and area2 is %0.2f\n", area1, area2)
}
