package main

import (
	"fmt"
)

func main() {
	result := beautifulDays(20, 23, 6)
	fmt.Println(result)
}

func beautifulDays(i int32, j int32, k int32) int32 {
	reversed := func(n int32) int32 {
		var rev int32 = 0
		for n != 0 {
			rev = (rev * 10) + n%10
			n /= 10
		}
		return rev
	}

	var beautyNumber int32
	for start := i; start <= j; start++ {
		r := (start - reversed(start)) % k
		if r == 0 {
			beautyNumber++
		}
	}

	return beautyNumber
}
