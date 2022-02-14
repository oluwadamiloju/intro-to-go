package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["first"] = 1
	x["second"] = 2

	delete(x, "first")
	fmt.Println(x)

	name, ok := x["Uno"]
	fmt.Println(name, ok)

	y := map[string]string{
		"firstName": "Sarah",
		"lastName":  "Akinkunmi",
		"age":       "Twenty-one",
	}

	fmt.Println(y["firstName"])
}
