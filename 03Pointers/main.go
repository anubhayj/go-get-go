package main

import "fmt"

func main() {

	//Declaring a pointer
	var ptr *int
	fmt.Println("Value of ptr, ", ptr)

	num := 23

	var ptr2 = &num
	fmt.Println("reference to  ptr2", ptr2)
	fmt.Println("valur of ptr2", *ptr2)

	//
}
