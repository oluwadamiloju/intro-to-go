package main

import "fmt"

func main() {
	numberList := []float64{1, 2, 3, 4, 5}
	fmt.Println(average(numberList))
}

func average(list []float64) float64 {
	total := 0.00

	for i := 0; i < len(list); i++ {
		total += list[i]
	}

	return total / float64(len(list))
}
