package main

import "fmt"

//There is no inheritence in golang
// No concept of super/parennt/child

type Subject struct {
	sub1 string
	sub2 string
}
type User struct {
	name    string
	class   int
	email   string
	subject Subject
}

func main() {

	Anubhay := User{"Anubhay", 5, "test@test.com", Subject{"Math", "CS"}}

	fmt.Println(Anubhay)
	fmt.Printf("Testing +V : %+v\n", Anubhay)

}
