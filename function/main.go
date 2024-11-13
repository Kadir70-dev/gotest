package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple function")
}
func add(a, b int) int {
	return a + b
}
func multiply(a, b int) (result int) {
    result = a * b
    return

}

func main() {
	fmt.Println("we are learning functions in golang")
	simpleFunction()

	ans := add(2, 3)
	fmt.Println("Add of two numbers is ", ans)
    fmt.Println("Multiply of two numbers is ", multiply(2, 3))
}
