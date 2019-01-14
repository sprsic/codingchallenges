package main

import (
	"fmt"
	"time"
)

/*
*Implement function takes in a function f
 and an integer n, and calls f after n milliseconds.
*/
func main() {
	delay(func(){	
		fmt.Printf("Hello\n");
	}, 5)
}

func delay(f func(), n int) {
	time.Sleep(time.Duration(n)* time.Second)
	f()
}

