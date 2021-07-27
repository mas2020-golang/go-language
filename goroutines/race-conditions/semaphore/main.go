/*
	This example shows how to avoid race simply using a semaphore. By this way the access
  to the common total variable is in sync.
*/
package main

import (
	"fmt"
	"time"
)

var total int
var sema = make(chan struct{},1)

func main() {
	// simulate some operations
	go deposit(10)
	go deposit(10)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("total is:", getTotal())
}

// deposit adds a number to the total using a channel
func deposit(n int) {
	sema <- struct{}{} // acquire the lock
	total += n
	<-sema // release the lock
}

// getTotal reads the total using a channel
func getTotal() int {
	return total
}
