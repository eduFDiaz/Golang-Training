// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	k := 2

	shiftedArray := shiftArray(x, k)
	fmt.Printf("tmp: %v, type of %T\n", shiftedArray, shiftedArray)
}

func shiftArray(array []int, k int) []int {
	tmp := append([]int{}, append(array[len(array)-k:len(array)], array[0:len(array)-k]...)...)
	return tmp
}
