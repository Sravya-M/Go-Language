package main

import "fmt"

func main() {

	num1, num2 := 2, 3

	// the () parantheses for condition are not necessary but no error on keeping them
	if num1 < num2 {
		fmt.Printf("%d is less than %d\n", num1, num2)
	}

	// the switch will not go to next case.. break is not required
	color := "red"

	switch color {
	case "blue":
		fmt.Println("blue color")
	case "red":
		fmt.Println("red color")
	default:
		fmt.Println("neither blue nor red")
	}

	// for loop
	// 1
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// 2
	for i := 1; i <= 2; i++ {
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 15; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
