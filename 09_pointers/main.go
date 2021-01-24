package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, &a)
	fmt.Println(*b, b)
	fmt.Printf("%T %T \n", a, b)

	// change values using pointers
	*b = 10
	fmt.Println(a)
}
