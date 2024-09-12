package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	counts := map[int32]int{}

	for _, bird := range arr {
		counts[bird]++
	}

	var lowerBird int32
	var maxVal int

	for bird, count := range counts {
		if count > maxVal {
			lowerBird = bird
			maxVal = count
		}
		if count == maxVal {
			if bird < lowerBird {
				lowerBird = bird
			}
		}
	}

	return lowerBird
}

func main() {
	n := []int32{1, 4, 4, 4, 5, 3}
	r := migratoryBirds(n)
	fmt.Println(r)
}
