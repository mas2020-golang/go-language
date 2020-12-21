package main

import (
	"fmt"
	"time"
)

/*
Main func is executed by the main goroutine.
Adding 'go' before call a func or a method causes the func is executed in a newly created
go routine.

*/
func main() {
	//fmt.Println(fib(6))
	// counter execution on a separate go routine
	go counter(1 * time.Second)
	// main go routine after X seconds will stop the application and all the goroutines abruptly terminated
	time.Sleep(3 * time.Second)
	fmt.Println()
}

/*
Execute an infinite loop and inside it a loop that shows a single char every delay time passed
as argument.
 */
func counter(delay time.Duration) {
	for {
		for _, r := range `1234567890` {
			fmt.Printf("\r%c/%d seconds", r,len(`1234567890`))
			time.Sleep(delay)
		}
	}
}

// Computes a Fibonacci number (1,1,2,3,5,8...)
func fib(x int) int {
	if x < 2 {
		fmt.Println("x:",x)
		return x
	}
	return fib(x-1) + fib(x-2)
}