package main

import "fmt"

func swap(a, b *int) { *a, *b = *b, *a }

func main() {
	var x, y int
	fmt.Print("Enter two integers: ")
	fmt.Scan(&x, &y)
	fmt.Printf("Before: x=%d, y=%d\nAfter: x=%d, y=%d\n", x, y, func(){ swap(&x, &y) }(); x, y)
}
