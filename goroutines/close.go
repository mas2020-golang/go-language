package main

import (
	"fmt"
	"time"
)

/*
	In this example we show how to close a channel.
*/

func main() {
	var c = make(chan int)
	go doSomething(c)
	for i := range c {
		fmt.Printf("value read from channel is %d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Main go routine end")
}

func doSomething(c chan int) {
	for i := 0; i < 2; i++ {
		if i == 0 {
			fmt.Println("#go routine I'm doing something...", i)
			time.Sleep(1 * time.Second)
			c <- i
			fmt.Println("#go routine value written on channel is ", i)
		}
		if i == 1 {
			fmt.Println("#go routine error close the channel...", i)
			time.Sleep(1 * time.Second)
			close(c)
			break
		}
	}
}
