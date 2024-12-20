package main

import (
	"fmt"
	"math/big"
)

// Fibonacci series function using Big Int to handle large numbers
func fibonacciBig(n int) []*big.Int {
	fib := make([]*big.Int, n)
	fib[0] = big.NewInt(0)
	fib[1] = big.NewInt(1)

	for i := 2; i < n; i++ {
		fib[i] = new(big.Int).Add(fib[i-1], fib[i-2])
	}
	return fib
}

// Fibonacci series using memoization
func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, found := memo[n]; found {
		return val
	}
	memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	return memo[n]
}

// Function for calculating Fibonacci using an iterative approach for better performance
func fibonacciIterative(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func main() {
	var n int
	fmt.Print("Enter the number of Fibonacci terms to print: ")
	_, err := fmt.Scan(&n)

	// Input validation to ensure a valid positive number
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid positive number!")
		return
	}

	// Using Big Int for large Fibonacci numbers
	result := fibonacciBig(n)
	fmt.Println("Fibonacci series using Big Int:")
	for _, v := range result {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Using memoization for optimized calculation
	memo := make(map[int]int)
	fmt.Println("Fibonacci series using Memoization:")
	for i := 0; i < n; i++ {
		fmt.Print(fibonacciMemo(i, memo), " ")
	}
	fmt.Println()

	// Using iterative approach for Fibonacci series (No concurrency)
	resultIterative := fibonacciIterative(n)
	fmt.Println("Fibonacci series using Iterative Approach:")
	for _, v := range resultIterative {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Asking for a specific Fibonacci number at an index
	var index int
	fmt.Print("Enter an index to get the Fibonacci number at that position: ")
	fmt.Scan(&index)

	// Check if index is within bounds
	if index <= 0 || index > n {
		fmt.Println("Please enter a valid index between 1 and", n)
		return
	}

	// Printing the Fibonacci number at the user-defined index
	fmt.Printf("Fibonacci number at index %d is: %d\n", index, fibonacciMemo(index-1, memo)) // Index is 1-based
}