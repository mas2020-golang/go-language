/*
	This example shows how to avoid race simply accessing to a variable only with a go routine.
	The other routines can add a number passing through a func and using the add channel.
  To get the total the routines can use a special channel for this purpose. The sum variable is not
  exposed so that the access to modify that values is done in sync, ensuring that the data are concurrency-safe.
*/
package main

import (
	"fmt"
	"time"
)

var add = make(chan int)   // send amount to total
var total = make(chan int) // receive the total

func main() {
	// starts the manager into another go routine
	go manager()
	// simulate some operations
	time.Sleep(1 * time.Second)
	go deposit(10)
	go deposit(10)
	fmt.Println("total is:", getTotal())
	time.Sleep(200 * time.Millisecond)
	go fmt.Println("total is:", getTotal())
	time.Sleep(500 * time.Millisecond)
}

// deposit adds a number to the total using a channel
func deposit(n int) {
	add <- n
}

// getTotal reads the total using a channel
func getTotal() int {
	return <-total // waits until something is available
}

func manager() {
	var sum int = 0
	for {
		fmt.Println("evaluating select...")
		select {
		case amount := <-add:
			sum += amount
			fmt.Println("received an amount of", amount)
		case total <- sum:
		}
	}
}

func printTotal() {
	fmt.Println("total is: ", total)
}
