package main

/*
	An example of the select statement: in this for cycle we create a buffered channel with one value of int.
  Then we have a for loop until 8 iterations starting from 0 to 7.
	- the first time the select will execute the statement: ch <- 0 and the for continues
  - now the buffered channel is full, so the write operation ch <- i is not possible
  - the first operation to pick a value from the channel is the only the select can execute
  So, due to the length of the buffered channel to 1, we print the odd value from the for loop.

	PAY ATTENTION: if we increase the length of the buffer the output is nondeterministic. When multiple cases
  are valid at the same time the select will pick one the them randomly.
*/

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 8; i++ {
		select {
		case x := <-ch:
			fmt.Println("read from the channel the value:", x) // "0" "2" "4" "6"
		case ch <- i:
			fmt.Println("write into the channel the value of:", i)
		}
	}
	fmt.Println("for loop is gone...bye")
}
