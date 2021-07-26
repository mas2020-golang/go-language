package main

import (
	"fmt"
	"time"
)

var total int = 10

func main() {
	go add(100)
	go add(200)
	time.Sleep(100 * time.Millisecond)
	printTotal()
	time.Sleep(100 * time.Millisecond)
}

func add (a int){
	total += a
}

func printTotal(){
	fmt.Println("total is: ", total)
}

