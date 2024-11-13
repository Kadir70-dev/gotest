package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	_, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Error handling in golang")
}
