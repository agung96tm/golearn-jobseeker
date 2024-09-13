package main

import "fmt"

func main() {
	result := pageCount(6, 2)
	fmt.Println(result)
}

func pageCount(n int32, p int32) int32 {
	fromFirstPage := p / 2
	fromLastPage := (n / 2) - fromFirstPage

	if fromFirstPage < fromLastPage {
		return fromFirstPage
	} else {
		return fromLastPage
	}
}
