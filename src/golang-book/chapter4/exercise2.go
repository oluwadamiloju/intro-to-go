package main

import "fmt"

func main() {
	fmt.Print("Enter distance in feet: ")
	
	var feet float64
	fmt.Scanf("%f", &feet)
	
	meters := feet * 0.3048
	fmt.Println(meters)
}