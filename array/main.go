package main

import "fmt"

func main() {
	fmt.Println("Arrays in golang")

	var names [5]string

	names[0] = "Kadir"
	names[3] = "Kadir"
	// names[2] = "Kadir"
	// names[3] = "Kadir"
	// names[4] = "Kadir"
	fmt.Println(names)
}