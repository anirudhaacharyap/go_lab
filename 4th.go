package main

import (
	"fmt"
	"math/big"
)

func factorialIter(n int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		res.Mul(res, big.NewInt(i))
	}
	return res
}

func factorialRec(n int64) *big.Int {
	if n < 2 {
		return big.NewInt(1)
	}
	return new(big.Int).Mul(big.NewInt(n), factorialRec(n-1))
}

func main() {
	var n int64
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("Factorial of %d (Iterative): %v\nFactorial of %d (Recursive): %v\n", n, factorialIter(n), n, factorialRec(n))
}
