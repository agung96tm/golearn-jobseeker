package main

import "fmt"

func main() {
	heights := []int32{1, 6, 3, 5, 2}
	var characterJump int32 = 4

	result := hurdleRace(characterJump, heights)
	fmt.Println(result)
}

func hurdleRace(k int32, height []int32) int32 {
	maxHeight := height[0]

	for _, h := range height[1:] {
		if h > maxHeight {
			maxHeight = h
		}
	}

	if maxHeight > k {
		return maxHeight - k
	}
	return 0
}
