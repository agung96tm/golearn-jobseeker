package main

import "fmt"

func main() {
	result := climbingLeaderboard(
		[]int32{100, 100, 50, 40, 40, 20, 10},
		[]int32{5, 25, 50, 120},
	)
	fmt.Println(result)
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	uniqueRanked := make([]int32, 0)
	for i := 0; i < len(ranked); i++ {
		if i == 0 || ranked[i] != ranked[i-1] {
			uniqueRanked = append(uniqueRanked, ranked[i])
		}
	}

	result := make([]int32, len(player))
	n := len(uniqueRanked)
	index := n - 1

	for i, score := range player {
		for index >= 0 && score >= uniqueRanked[index] {
			index--
		}
		result[i] = int32(index + 2)
	}

	return result
}
