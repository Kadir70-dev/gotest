package main

import "fmt"

func main() {
	fmt.Println("Switch in golang")
	color := "red"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}
}


