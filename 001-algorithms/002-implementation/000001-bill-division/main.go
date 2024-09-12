package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var total int32
	for itemKey, itemVal := range bill {
		if itemKey == int(k) {
			continue
		}
		total += itemVal
	}

	charged := b - (total / 2)
	if charged == 0 {
		fmt.Print("Bon Appetit")
	} else {
		fmt.Printf("%d\n", charged)
	}
}

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 1, 12)
}
