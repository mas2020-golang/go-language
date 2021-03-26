package main

import (
	"fmt"
	"os"
	"time"
)

/*
	Example on how to using multiplexing with select. The select statement is a way to wait for completion from more
	channels.
	Think the case we have two channels and we need to catch the value from each other. The first we listen will block
	eventually a value received on the other. The way to accomplish that is using a channel. Let's have a look.
	A select waits until a communication for some case is ready to proceed.
	It then performs that case executing the associated statements; the other communications do not happen.
	Using the `default` clause means that specify what to do if none of the others cases can proceed immediately.
	Usually it's used for polling.
*/

func main() {
	//countdown()
	//countdownTicker()
	waitForLaunch()
	panic(nil)
}

func waitForLaunch() {
	abort := make(chan struct{})
	// Separate go routine listening that sends an empty structure on the abort channel in
	// case the user will press enter
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	go count(time.Second * 1)

	// this code is executed on the main go routine and will block until a case is not satisfied: in that case
	// it will execute the corresponding statements and the will go on to the following statements (if any) after
	// the select
	fmt.Println("Press return to abort (you have three seconds)...")
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("\nFive seconds lasted...go on without interruptions")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Launch take time, please wait 3 seconds...")
	time.Sleep(3 * time.Second)
	fmt.Println("Launch is gone!")
}

func count(delay time.Duration) {
	i := 0
	for {
		fmt.Printf("\r%d seconds", i)
		time.Sleep(delay)
		i++
	}
}
