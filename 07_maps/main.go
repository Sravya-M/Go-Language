package main

import "fmt"

func main() {

	// declaration
	students := make(map[int]string)

	// assigning key value pairs
	students[1] = "sravya"
	students[10] = "swati"

	fmt.Println(students)
	fmt.Println(students[1])

	delete(students, 10)

	// assignment and declaration together
	studentsNew := map[int]string{1: "sravya", 2: "swati"}
	fmt.Println(studentsNew)

}
