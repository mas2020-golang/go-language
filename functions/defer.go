package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	defer traceTime("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(1 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

/*
Use defer to register start, end and total time of a func
 */
func traceTime(funcName string) func() {
	start := time.Now()
	fmt.Fprintf(os.Stdout, "%s [start] %s\n", start.Format("2006-01-02T15:04:05.000"), funcName)
	return func() {
		end := time.Now()
		elapsedTime := end.Sub(start)
		fmt.Fprintf(os.Stdout, "%s [end] %s (%1.3f seconds)\n", end.Format("2006-01-02T15:04:05.000"), funcName, elapsedTime.Seconds())
	}
}

