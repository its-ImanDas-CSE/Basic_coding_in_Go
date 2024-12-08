package main

import (
	"fmt"
)

// Recursive function to calculate Fibonacci number
func fibonacci_recursion(Num int) int {
	if Num == 0 {
		return 0 // Fibonacci of 0 is 0
	} else if Num == 1 {
		return 1 // Fibonacci of 1 is 1
	} else {
		return fibonacci_recursion(Num-1) + fibonacci_recursion(Num-2)
	}
}

func main() {
	fmt.Println("Give the number of terms in the Fibonacci series you want:")
	var Num int
	fmt.Scanln(&Num)

	// Print the Fibonacci series up to the Num-th term
	for i := 0; i < Num; i++ {
		fmt.Print(fibonacci_recursion(i), " ")
	}
}
