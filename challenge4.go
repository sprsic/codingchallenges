package main

import (
	"fmt"
	"math"
)

/*

Given an array of integers, find the first missing positive integer in linear time and constant space.
In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.
*/

func main() {

	array := []int{-1, 1, -2, -3, -4, 99, 4}
	fmt.Printf("%v", findMissingNumber(array))
}

func findMissingNumber(array []int) int {
	l := len(array)
	FLAG := l + 2

	for i, num := range array {
		if num <= 0 || num > len(array) {
			array[i] = FLAG
		}
	}

	// missing number is in range of 0 - len(array)+1
	// if number n = [0, len(array)+1] is in array than put number on position n negative
	for _, num := range array {
		if num == FLAG {
			continue
		}
		//sanitize negative
		val := int(math.Abs(float64(num)))
		//val - 1 because array begins from 0
		array[val-1] = -int(math.Abs(float64(array[val-1])))
	}
	// index where numner is positite is the one that is missing
	for i, num := range array {
		if num > 0 {
			return i + 1
		}
	}

	return l + 1
}
