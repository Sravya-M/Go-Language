package main

import "fmt"

// Inheritance is not supported in Go as there are no classes
// Used Composition, where struct is used to form other objects
// Multiple Inheritance is also suuported

type first struct {
	base1 string
}

type second struct {
	base2 string
}

func (f first) print() string {
	return f.base1
}

func (f second) display() string {
	return f.base2
}

type child struct {
	// struct embedding
	first
	second
}

func main() {
	c := child{first{base1: "base struct 1"},
		second{base2: "base struct 2"},
	}
	fmt.Println(c.print(), c.display())
}
