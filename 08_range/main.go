package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	// id is each item in ids
	for _, id := range ids {
		fmt.Println(id)
	}

	// i is the index starting from 0
	names := []string{"sravya", "kavya"}
	for i, id := range names {
		fmt.Println(i, id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// for range in maps
	students := map[int]string{1: "sravya", 2: "kavya"}
	for k, v := range students {
		fmt.Println(k, v)
	}
}
