package main

import "fmt"

func main() {
	n := Fibonacci(1)
	fmt.Println(n)
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	prev1, prev2 := 1, 1
	currentVal := 0
	for i := 3; i <= n; i++ {
		currentVal = prev2 + prev1
		prev1 = prev2
		prev2 = currentVal
	}

	return currentVal
}
