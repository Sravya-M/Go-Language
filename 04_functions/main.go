package main

import "fmt"

// see the arguments in add, sub functions.. as both are int we can write (num1, num2 int)
func add(num1 int, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 - num2
}

func main() {
	fmt.Println("Hi")
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
}
