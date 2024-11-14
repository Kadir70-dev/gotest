package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	studentGrades := make(map[string]int)
	studentGrades["Kadir"] = 100
	studentGrades["Mukadam"] = 90
	studentGrades["raza"] = 80
	fmt.Println(studentGrades)
	fmt.Println(studentGrades["Kadir"])

}
