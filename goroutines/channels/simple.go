/*
Through the channels is possible for a go routine send values to another. The channel is a conduit of a
particular type.
'chan int' means a channel of an int type as element

A channel is a reference to the data structure created by make.
Copying a channel or pass one as an argument to a function, the reference is passed: an update will reflect on the caller
and callee as they refer to the same data structure.
* The zero value of a channel is nil.
* comparison between two channel is true if they refer to the same data structure

To send a value through the channel: channel <- value operands
To receive a value from a channel: <-channel operand

c := make(chan string)
c <- "hello" // to send
x := <-c // to receive
<-c // receive discarding result

To close a channel use close(channel):
- send data on a closed channel will panic
- receive on a close channel yields values until no more values are left

There are two types of channel:
- 0 capacity, called `unbuffered`
- with capacity, called `buffered`
You create the corresponding type passing the capacity arg to the make function.

IMPORTANT:
- send operation on unbuffered channel (with zero capacity) blocks the sending goroutine until another
goroutine receive the data. From this point both goroutine can continue.
- if the receive operation was attempted first, the receiving goroutine is locked until another doesn't send data
- both goroutines on unbuffered channel are synchronous.
- when a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the
sending goroutine
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	// invoke single job as go routine
	go singleJob(done)
	fmt.Println("I am waiting the end of singleJob...")
	// lock the routine until a message is read from the done channel
	<-done
	fmt.Println("I can terminate the application, bye in 3 sec...")
	go counter()
	time.Sleep(3 * time.Second)
}

// Single job as a go routine
func singleJob(done chan struct{}) {
	fmt.Println("singleJob => I am doing a complex job...")
	time.Sleep(2 * time.Second)
	// write on the channel
	done <- struct{}{}
	fmt.Println("singleJob => read done from other routines")
}

// Simple seconds time counter
func counter() {
	i := 0
	for {
		fmt.Printf("\r%d seconds...", i)
		time.Sleep(1 * time.Second)
		i++
	}
}
