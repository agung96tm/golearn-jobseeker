package main

import "fmt"

func main() {
	design := designerPdfViewer(
		[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		"abc",
	)

	fmt.Println(design)
}

func designerPdfViewer(h []int32, word string) int32 {
	var tallest int32
	for _, w := range word {
		// Get height of the current letter by using its position in the alphabet
		height := h[w-'a']

		if height > tallest {
			tallest = height
		}
	}

	return tallest * int32(len(word))
}
