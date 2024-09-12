package main

import "fmt"

func main() {
	result := sockMerchant([]int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	fmt.Println(result)
}

func sockMerchant(ar []int32) int32 {
	counter := make(map[int32]int)

	for _, sock := range ar {
		counter[sock]++
	}

	var total int32
	for _, sock := range counter {
		total += int32(sock / 2)
	}
	return total
}
