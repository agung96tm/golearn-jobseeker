package main

import "fmt"

func main() {
	result := viralAdvertising(5)
	fmt.Println(result)
}

func viralAdvertising(n int32) int32 {
	shared := int32(5)
	cumulative := int32(0)

	for i := int32(0); i < n; i++ {
		liked := shared / 2
		shared = liked * 3
		cumulative += liked
	}

	return cumulative
}
