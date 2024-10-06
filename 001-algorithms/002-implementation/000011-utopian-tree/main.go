package main

import "fmt"

func main() {
	growthCycles := []int32{4}

	for _, n := range growthCycles {
		result := utopianTree(n)
		fmt.Println(result)
	}
}

func utopianTree(n int32) int32 {
	var height int32 = 1

	for i := int32(1); i <= n; i++ {
		if i%2 == 1 {
			height *= 2
		} else {
			height += 1
		}
	}

	return height
}
