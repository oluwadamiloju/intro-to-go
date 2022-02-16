package main

import "fmt"

func multiply(args ...float64) float64 {
	total := 1.0
	for _, i := range args {
		total *= i
	}

	return total
}

func main() {
	fmt.Println(multiply(1, 2, 3, 4, 5))
	fmt.Println(multiply(4, 2, 4, 1))
}
