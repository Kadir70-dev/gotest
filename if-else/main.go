package main

import "fmt"


func main() {
	x:= 2
	if x > 5{
      fmt.Println("x is greater than 5")
	}  else {
		fmt.Println("x is smaller than 5")
	}

	z:= 10
	if z > 5{
		fmt.Println("z is greater than 5")
	} else if z == 5 {
		fmt.Println("z is equal to 5")
	} else {
		fmt.Println("z is smaller than 5")
	}
}