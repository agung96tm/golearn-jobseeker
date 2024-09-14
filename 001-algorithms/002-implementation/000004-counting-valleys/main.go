package main

import "fmt"

func countingValleys(_ int32, path string) int32 {
	walkCounter := 0
	var valleyCounter int32 = 0
	for _, p := range path {
		step := string(p)
		if step == "U" {
			walkCounter++
			if walkCounter == 0 {
				valleyCounter++
			}
		} else {
			walkCounter--
		}
	}
	return valleyCounter
}

func main() {
	result := countingValleys(8, "DDUUUUDD")
	fmt.Println(result)
}
