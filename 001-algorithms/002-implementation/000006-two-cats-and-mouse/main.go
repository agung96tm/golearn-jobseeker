package main

import "fmt"

func catAndMouse(catA, catB, mouseC int32) string {
	var abs = func(r int32) int32 {
		if r < 0 {
			return -r
		}
		return r
	}

	aDistance := abs(catA - mouseC)
	bDistance := abs(catB - mouseC)

	switch {
	case aDistance < bDistance:
		return "Cat A"
	case bDistance < aDistance:
		return "Cat B"
	default:
		return "Mouse C"
	}
}

func main() {
	var a int32 = 2
	var b int32 = 5
	var c int32 = 4

	fmt.Println(catAndMouse(a, b, c))
}
