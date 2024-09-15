package main

import "fmt"

func main() {
	keyboards := []int32{40, 50, 60}
	drives := []int32{5, 8, 12}
	var budget int32 = 60

	result := getMoneySpent(keyboards, drives, budget)
	fmt.Println(result)
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maxSpent int32 = -1

	for _, keyboard := range keyboards {
		for _, drive := range drives {
			total := keyboard + drive

			if total > maxSpent && total <= b {
				maxSpent = total
			}
		}
	}

	return maxSpent
}
