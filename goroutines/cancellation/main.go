package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
	This example shows how to cancel a go routine that is working using a done channel that is checked every
	time the go routine starts a new work cycle.
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
func doSomething(routine int, ch chan<- int) {
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
		fmt.Println("cancelled is done!")
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
		go doSomething(i, listen)
	}

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// wait for the go routine to end before exit
loop:
	for i := 0; i < 3; i++ {
		select {
		case routine, ok := <-listen:
			if ok {
				fmt.Printf("receive end signal from go routine #%d\n", routine)
			} else {
				break loop
			}
		case <-time.After(10 * time.Second):
			fmt.Println("timeout exceeded, application will be forced to stop...")
			close(done)
			time.Sleep(5 * time.Second) // some time to close other go routine (not deterministic)
			break loop
		}
	}
	fmt.Println("Main go routine end...")
}
