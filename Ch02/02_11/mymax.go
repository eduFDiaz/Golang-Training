package main

import (
	"fmt"
)

func main() {
	array := []int{16, 8, 42, 4, 23, 15}
	fmt.Printf("max val=%v", max(array))
}

func max(array []int) int {
	max := array[0]
	for _, element := range array[1:] {
		if element > max {
			max = element
		}
	}
	return max
}
