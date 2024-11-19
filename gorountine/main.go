package main

import (
	"fmt"
	"time"
)



func sayHello() {
	fmt.Println("1")
	time.Sleep(2000*time.Millisecond) //simulating some work
	fmt.Println("2")
}

func sayHi() {
	fmt.Println("3")
}
func main() {
	fmt.Println("4")
    go	sayHello()
	go sayHi()

	time.Sleep(3000*time.Millisecond)
}