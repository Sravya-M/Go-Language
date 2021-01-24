package main

import "fmt"

// gives error if we use the below
//global := "shorthand assignment works only in functions"

func main() {
	fmt.Println("Hello")

	// we need not specify type for a string, it infers
	var name string = "srav"
	var gender = "Female"
	fmt.Println(name, gender)

	// to print type
	fmt.Printf("%T\n", name)

	// you will get an error if a variable is not used but declared
	//	var isCool = true

	local := "shorthand assignment works only in functions"
	fmt.Println(local)

}
