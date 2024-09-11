package main

import "fmt"

func main() {
	items := []int{5, 4, 1, 2, 99}

	BubbleSort(items)
	fmt.Println(items)
}

func BubbleSort(items []int) {
	if len(items) <= 1 {
		return
	}

	for {
		isSorted := false

		for i := 0; i < len(items)-1; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				isSorted = true
			}
		}

		if isSorted == false {
			break
		}
	}
}
