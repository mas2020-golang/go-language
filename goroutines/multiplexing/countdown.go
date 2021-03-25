package main

import (
	"fmt"
	"os"
	"time"
)

/*
	Example of a countdown using the Tick func of the time package.
*/

func countdown() {
	abort := make(chan struct{})
	// Separate go routine listening that sends an empty structure on the abort channel in
	// case the user will press enter
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	// this code is executed on the main go routine and will block until a case is not satisfied: in that case
	// it will execute the corresponding statements and the will go on to the following statements (if any) after
	// the select
	fmt.Println("Press return to abort...")

	/*
		every one sec a Time object is sent to the tick <- chan time.Time. PAY ATTENTION: tick is convenient but is
		useful only when the tick is needed throughout the lifetime of the application because the ticker go routine is
		still there trying to send values in a channel from which nobody is receiving, the risk is a go routine leak.
	*/
	tick := time.Tick(1 * time.Second)
	for c := 9; c > 0; c-- {
		select {
		case <-tick:
			//fmt.Println("time is", t)
			fmt.Printf("\r%d", c)
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	println()
	launch()
}

// Countdown using the Ticker
func countdownTicker() {
	abort := make(chan struct{})
	// Separate go routine listening that sends an empty structure on the abort channel in
	// case the user will press enter
	go func() {
		fmt.Println("Press return to abort...")
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		fmt.Println("return pressed, aborting...")
		abort <- struct{}{}
	}()

	// this code is executed on the main go routine and will block until a case is not satisfied: in that case
	// it will execute the corresponding statements and the will go on to the following statements (if any) after
	// the select

	/*
		using the Ticker prevent from go routine leak because using the Stop() method causes the ticker's go routine
		to terminate
	*/
	ticker := time.NewTicker(1 * time.Second)
	for c := 9; c > 0; c-- {
		select {
		case <-ticker.C:
			//fmt.Println("time is", t)
			fmt.Printf("\r%d", c)
		case <-abort:
			fmt.Println("Launch aborted!")
			ticker.Stop()
			time.Sleep(500 * time.Millisecond) // to take time go routine to close (not necessary, only for test purpose)
			return
		}
	}
	ticker.Stop()
	println()
	launch()
}
