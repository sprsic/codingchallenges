package main

import (
	"fmt"
)

/*Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.*/
func main() {
	array := []int{10, 15, 3, 7}
	sumK(17, array)
}

func sumK(sum int, array []int) bool {

	m := make(map[int]struct{})

	for _, num := range array {
		m[num] = struct{}{}
	}

	var s = array[0:len(array)]
	for _, n := range s {
		sum = sum - n
		if m[sum] == struct{}{} {
			fmt.Printf("{ %d, %d }", n, sum)
			return true
		}
	}

	return false
}
