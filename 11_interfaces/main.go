package main

import (
	"fmt"
	"math"
)

// Shape interfaces contain method signatures for structs
type Shape interface {
	area() float64
}

// Circle structure
type Circle struct {
	radius float64
}

// Rectangle structure
type Rectangle struct {
	width, height float64
}

// test for github creds
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Rectangle) area() float64 {
	return c.width * c.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{5}
	rectangle := Rectangle{5, 6}
	fmt.Println(getArea(circle), getArea(rectangle))
}
