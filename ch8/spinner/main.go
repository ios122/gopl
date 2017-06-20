// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 218.

// Spinner displays an animation while computing the 45th Fibonacci number.
package main

import (
	"fmt"
	"time"
)

//!+
func main() {
	//go spinner(100 * time.Millisecond)
	//const n = 45
	//fibN := fib(n) // slow
	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)


	var timeNow = time.Now()

	utcHour := timeNow.In(time.UTC).Hour()

	diffHour := (10 - utcHour)

	if diffHour <= -12 {
		diffHour += 24
	}
	ret := []string{fmt.Sprintf("%+.4d", diffHour*100), fmt.Sprintf("%+.4d", diffHour*100+30)}

	fmt.Println(ret)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//!-
