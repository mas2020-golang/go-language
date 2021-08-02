/*
	This example shows how to avoid race simply using the sync package. This example is the evolution
	of using a semaphore directly
*/
package main

import (
  "fmt"
  "sync"
  "time"
)

var mu sync.Mutex
var total int

func main() {
  // simulate some operations
  go deposit(10)
  go deposit(10)
  time.Sleep(200 * time.Millisecond)
  fmt.Println("total is:", getTotal())
}

// deposit adds a number to the total using a channel
func deposit(n int) {
  mu.Lock() // acquire the lock
  total += n
  mu.Unlock() // release the lock
}

// getTotal reads the total using a channel
func getTotal() int {
  mu.Lock() // acquire the lock
  defer mu.Unlock() // release the lock
  return total
}
