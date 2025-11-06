package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Square struct{ side float64 }
func (s Square) area() float64       { return s.side * s.side }
func (s Square) perimeter() float64  { return 4 * s.side }

type Circle struct{ radius float64 }
func (c Circle) area() float64       { return math.Pi * c.radius * c.radius }
func (c Circle) perimeter() float64  { return 2 * math.Pi * c.radius }

func main() {
	s, c := Square{5}, Circle{3}
	fmt.Printf("Square: Area=%.2f, Perimeter=%.2f\n", s.area(), s.perimeter())
	fmt.Printf("Circle: Area=%.2f, Perimeter=%.2f\n", c.area(), c.perimeter())
}
