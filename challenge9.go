package main

import (
    "fmt"
    "math"
)

/*
Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.

*/
func main() {

    array:= [] int {2, 4, 6, 2, 5}
    sum:= nonAdjacent(array)
    fmt.Printf("%d\n", sum)
}

func nonAdjacent(array[] int) int {
    inclusive:= array[0]
    exclusive:= 0
    length:= len(array)
    pom:= exclusive
    for i:= 1;i < length; i++ {

        if inclusive > exclusive {
            pom = inclusive
        } else {
            pom = exclusive
        }

        inclusive = array[i] + exclusive
        exclusive = int(math.Max(float64(exclusive), float64(pom)))

    }

    return int(math.Max(float64(inclusive), float64(exclusive)))
}