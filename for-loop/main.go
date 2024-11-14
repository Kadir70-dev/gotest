package main

import (
	"fmt"
)

func main() {
	fmt.Println("for loop in golang")
	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}

	counter := 0
	for {
		fmt.Println("infinite loop")
		counter++
		if counter == 1 {
			break
		}
	}

	number := []int{1, 2, 3, 4, 5}
	for index, value := range number {
		fmt.Println(index, value)
	}

}
