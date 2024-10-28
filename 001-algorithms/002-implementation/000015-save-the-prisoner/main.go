package main

import "fmt"

func main() {
	result := saveThePrisoner(4, 6, 2)
	fmt.Println(result)
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	result := (s + m - 1) % n
	if result == 0 {
		return n
	}
	return result
}
