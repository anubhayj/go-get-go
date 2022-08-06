package main

import "fmt"

func main() {

	var fruit [5]string

	fruit[0] = "Apple"
	fruit[1] = "Guava"
	fruit[4] = "Grapes"

	fmt.Println(" Fuits :", fruit)
	fmt.Println(" Fuits :", len(fruit)) //Since you have reserved memory, it will always give the len that you declared

	var list2 = [3]string{"1", "2", "3"}

	fmt.Println(" veg :", len(list2))

	/* Slice */
	// Widely used in golang
	// especially with DataBases

	var myslice = []string{"1", "2"}
	fmt.Println(" myslice :", len(myslice))
	fmt.Printf(" type of myslice %T\n :", myslice)

	//Adding to slice
	myslice = append(myslice, "4", "5")
	fmt.Printf(" type of myslice : %T\n", myslice)

	// slice[x,y]  the x and y are non inclusive

	/*
		make will initialise the elements with 0 [0,0,0,0]
		make will give default allocation of memory, but you can increase the size with append
	*/
	var arr = make([]int, 4)
	fmt.Println("arr", arr)

	//Delete from slice

	var num = []string{"1", "2", "3", "4", "5"}
	var indx int = 2

	num = append(num[:indx], num[indx+1:]...)
	fmt.Println(num)

}
