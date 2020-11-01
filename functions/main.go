package main

import "fmt"

/*
	This example shows how to use functions in Go.
*/

func main() {
	// call a function with two params that comes back an int
	retValue := example(100, "test")
	fmt.Println("retValue is", retValue)
	fmt.Printf("Type of example is %T\n", example)

	// call a function that returns multiple values
	ret1, ret2 := multipleReturns(100, "test")
	fmt.Printf("ret1: %d, ret2: %s\n", ret1, ret2)
	fmt.Printf("Type of exampleNoReturn is %T\n", exampleNoReturn)

	// call a function that returns multiple named values
	ret1, ret2 = multipleNamedReturns(200, "test")
	fmt.Printf("---multipleNamedReturns ret1: %d, ret2: %s\n", ret1, ret2)

	// recursive example
	searchHtml()
}

/*
	This func has two params (int and string) and returns to the
	caller a int value.
*/
func example(par1 int, par2 string) int {
	fmt.Println("the par2 has this value: ", par2)
	return par1 * 2
}

func exampleNoReturn(par1 int, par2 string) {
	fmt.Println("the par2 has this value: ", par2)
}

/*
	This func has two params (int and string) and returns to result to the
	caller.
*/
func multipleReturns(par1 int, par2 string) (int, string) {
	fmt.Println("the par2 has this value: ", par2)
	return par1 * 2, par2 + "-return"
}

/*
	This func has two params (int and string) and returns to result to the
	caller.
*/
func multipleNamedReturns(par1 int, par2 string) (val int, name string) {
	fmt.Println("the par2 has this value: ", par2)
	val = par1 * 2
	name = par2 + "-test"
	return
}