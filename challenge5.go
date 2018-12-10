package main

import "fmt"

/**
cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

Given this implementation of cons:
*/

type  pair struct {
	a, b int
}

func main() {
    pair := cons(3,4)
    fmt.Printf("%d \n",car(pair))
    fmt.Printf("%d \n",cdr(pair))
}

func car(pair pair) int {
    return pair.a
}

func cdr(pair pair) int {
    return pair.b
}

func cons(a int, b int) pair {
	return pair{a,b}
}
