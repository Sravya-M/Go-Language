package main

import (
	"fmt"
	"strconv"
)

// Student struct
type Student struct {
	age  int
	name string
	// if we have more variables with same type, declare in 1 line (fname, lname, name string)
}

// value receiver (no change is done on values)
// func (struct object) funcName(arguments) returnType
func (s Student) greet(msg string) string {
	// the s here is just a copy not the exact location
	s.age--
	return msg + " I am " + s.name + ", " + strconv.Itoa(s.age)
}

// pointer receiver
func (s *Student) incrAge() {
	s.age++
}

func main() {

	// 1
	student1 := Student{age: 25, name: "kavya"}

	// 2
	student2 := Student{25, "sravya"}

	student1.age++
	fmt.Println(student1.greet("Hello"))

	fmt.Println(student1)

	student2.incrAge()
	fmt.Println(student2)
}
