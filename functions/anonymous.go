package main

import "fmt"

/*
This func returns an anonymous func that stores the i variable as a value. By this way
the i variable has a scope that goes over of counter.
 */
func counter() func() int{
	var i int
	return func() int {
		i++
		return i
	}
}

func useCounter(){
	f := counter() // f is func() of type func() int
	fmt.Printf("call anonymous func, value is %d\n", f())
	fmt.Printf("call anonymous func, value is %d\n", f())
	fmt.Printf("call anonymous func, value is %d\n", f())
}

