package main

import "fmt"

func main() {
	fmt.Print("Enter the temperature in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	c := ((fahrenheit - 32) * 5) / 9
	fmt.Println(c)
}