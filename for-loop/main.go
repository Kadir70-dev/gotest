package main

import (
	"fmt"
)

func main() {
	fmt.Println("for loop in golang")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter := 0
	for {
		fmt.Println("infinite loop")
		counter++
	}
}
