package main

import "fmt"

func main() {
	// Arrays
	// 1.
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Örange"

	fmt.Println(fruitArr)

	// 2. declare and assign values
	fruitArr2 := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr2)

	// 3
	intArr := [...]int{1, 2, 3}
	fmt.Println(intArr)
	fmt.Println(len(intArr))

	// Slice - the number of elements is not fixed
	fruitSlice := []string{"Apple", "Orange", "Grape", "Guava"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
