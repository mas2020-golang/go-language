package main

import (
	"fmt"
	"time"
)

/*
This example shows how to manage types in Go.
*/

func main() {
	// Check for a type assertion
	typeAssertion(12)
	typeAssertion(map[string]interface{}{
		"key":  12,
		"key2": 13,
	})
}

func typeConversion() {
	/*
		This example convert some types into others
	*/
	var i int = 100
	var f float64 = float64(i) // convert int => float64
	var u uint = uint(f)       // convert float64 to an unsigned int
	fmt.Println("u is ", u)
}

func typeAssertion(y interface{}) {
	/*
		ok is a boolean that is true if the type corresponds to the type passed. If true z
		will contain the value of the type, else the zero value for that type.
	*/
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
		fmt.Println(z)
	} else {
		fmt.Println("y is not of the type map[string]interface{}")
	}

	// check a type with switch..case
	switch str := y.(type) {
	case int:
		fmt.Println("int type passed as argument:", str)
	case map[string]interface{}:
		fmt.Println("y is map[string]interface{}:", str)
	}
}
