package main

import "fmt"

/*
	Channels can be used to connect goroutines together so that the output of one is the input to another.
	In the example there are 3 go routines in order thanks to the channel sync.
*/

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			fmt.Printf("write %d into naturals\n", x)
			naturals <- x
			// it is possible to close a channel
			if x == 10 {
				close(naturals)
				break
			}
		}
	}()

	// Squarer
	go func() {
		for {
			fmt.Println("wait to read from naturals")
			// check first if the channel is closed or not
			x, ok := <-naturals
			if !ok {
				break
			}
			fmt.Printf("read %d from naturals\n", x)
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
