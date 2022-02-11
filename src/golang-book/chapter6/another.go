package main

import "fmt"

func main() {
	// first method of creating array

	//var x [5]float64
	//x[0] = 98
	//x[1] = 93
	//x[2] = 77
	//x[3] = 82
	//x[4] = 83

	// second method
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0

	for _, value := range x { //instead of {i < len(x); i++}, use { variable name :=  range of variable you want to loop over}
		total += value
	}

	fmt.Println(total / float64(len(x)))
}
