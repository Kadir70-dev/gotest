package main

import "fmt"


type Person struct {
	name string
	age int
}

func main() {
	fmt.Println("struct in golang")
	var prince Person
	prince.name = "Kadir"
	prince.age = 25
	fmt.Println(prince)
	fmt.Println(prince.name)
	fmt.Println(prince.age)
}