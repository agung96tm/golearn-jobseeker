package main

import "fmt"

func main() {
	status := angryProfessor(2, []int32{-1, -3, 4, 2})
	fmt.Println(status)
}

func angryProfessor(k int32, a []int32) string {
	var total int32 = 0
	for _, i := range a {
		if i <= 0 {
			total++
		}
		if total >= k {
			return "NO"
		}
	}
	return "YES"
}
