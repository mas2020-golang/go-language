package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
	This example shows how to cancel a go routine that is working.
*/

// listen channel (nil)
var listen = make(chan int)
var done chan bool = make(chan bool)

func main() {
	withCancellation()
}

/*
	This routine use the select statement for multiplexing listening from two different channels.
	In case the done channel receive a value the routine will be stopped.
*/
func doSomething(routine int, done chan bool, ch chan<- int) {
	fmt.Printf("#%d go routine start\n", routine)

	// do some work
	for i := 0; i < 3; i++ {
		if cancelled() {
			fmt.Printf("#%d go routine cancelled\n", routine)
			break
		}
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		seconds := r1.Intn(10)
		fmt.Printf("#%d go routine is working (%d seconds, iteration %d)\n", routine, seconds, i)
		time.Sleep(time.Duration(seconds) * time.Second)
	}

	// write the response into the channel
	ch <- routine
	fmt.Printf("#%d go routine end\n", routine)
}

// Return true if the go routines must stopped
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

/*
  Using the closing of the channel some routines can close and return but there is a problem.
	Not all the routine will be closed.
*/
func withCancellation() {
	fmt.Println("Main go routine start...")
	for i := 0; i < 3; i++ {
		go doSomething(i, done, listen)
	}

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// close the channel (closing a channel means to drain the values and then to get forever the zero value of the channel type)
	//close(done)
	// wait for the go routine to close before end
loop:
	for i := 0; i < 3; i++ {
		select {
		//case <-done:
		//	break loop
		case routine, ok := <-listen:
			if ok {
				fmt.Printf("receive end from routine #%d\n", routine)
			} else {
				break loop
			}
		}
	}
	fmt.Println("Main go routine end...")
	time.Sleep(200 * time.Millisecond)
}
