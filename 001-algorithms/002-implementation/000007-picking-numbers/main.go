package main

import "fmt"

func pickingNumbers(arr []int32) int32 {
	freq := make(map[int32]int32)
	for _, num := range arr {
		freq[num]++
	}

	maxLength := int32(0)

	for num := range freq {
		currentLength := freq[num] + freq[num+1]
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	result := pickingNumbers([]int32{4, 6, 5, 3, 3, 1})
	fmt.Println(result)
}
