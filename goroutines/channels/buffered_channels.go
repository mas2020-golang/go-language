package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main routine start...")
	fmt.Println("Mirrored query = ", mirroredQuery())
	time.Sleep(2 * time.Second)
	fmt.Println("Main routine end")
}

/*
	Using a buffered channel permits to end the three go routines independently
	from who ends first.
*/
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("asia.gopl.io")
		fmt.Println("end go routine asia.gopl.io")
	}()
	go func() {
		responses <- request("europe.gopl.io")
		fmt.Println("end go routine europe.gopl.io")
	}()
	go func() {
		responses <- request("americas.gopl.io")
		fmt.Println("end go routine americas.gopl.io")
	}()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	/* ... */
	fmt.Println("check hostname", hostname)
	time.Sleep(2 * time.Second)
	return hostname + " 100"
}
