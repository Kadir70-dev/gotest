package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice in golang")
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
