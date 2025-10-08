package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Please enter a positive number.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Print(fib(i), " ")
	}
}
