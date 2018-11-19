package main

import (
	"fmt"
)

/* Given an array of integers, return a new array such that each element at index i of the
new array is the product of all the numbers in the original array except the one at i.
 For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
  If our input was [3, 2, 1], the expected output would be [2, 3, 6].*/
func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v", productNaive(array))
	fmt.Printf("%v", productWithDivide(array))

}

// O(n^2) complexity
func productNaive(array []int) []int {
	len := len(array)
	var productArray = make([]int, len)

	var index int
	for i := range array {
		index = i
		product := 1
		for j, n := range array {
			if j != index {
				product = product * n
			}
		}
		productArray[i] = product
	}
	return productArray
}

// complexity O(n)
func productWithDivide(array []int) []int {
	product := 1

	for _, num := range array {
		product = product * num
	}

	productArray := make([]int, len(array))
	for i, num := range array {
		productArray[i] = product / num
	}

	return productArray
}
