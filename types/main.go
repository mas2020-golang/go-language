package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
This example shows how to manage types in Go.
*/
type Meters float64
type Miles float64

func main() {
	// Check for a type assertion
	typeAssertion(12)
	typeAssertion(map[string]interface{}{
		"key":  12,
		"key2": 13,
	})

	// Check the type conversion
	typeConversion()
}

func typeConversion() {
	/*
		This example converts some types into others
	*/
	var i, i2 int64 = 100, 0
	var f float64 = float64(i) // convert int => float64
	var u uint = uint(f)       // convert float64 to an unsigned int
	// to convert from string to boolean, int you need the strconv package
	b, _ := strconv.ParseBool("true")
	// convert 1000 in base 10
	i, _ = strconv.ParseInt("1000", 10, 0)
	i2, _ = strconv.ParseInt("010101101", 2, 0)
	fmt.Println("u is ", u)
	fmt.Println("bool is ", b)
	fmt.Println("int64 from 1000 is ", i) // => 1000
	fmt.Println("int64 010101101 is ", i2) // => 173

	var meters Meters = 1000
	var miles Miles = Miles(meters) * Miles(0.621371)
	var c int = int(miles) // value will be 621
	fmt.Println("meters, miles, c are:", meters, b, c)

	// Comparisons
	fmt.Println("meters == 0", meters == 0)
	fmt.Println("meters > 0", meters > 0)
	//fmt.Println("meters == miles (Meter == Miles)", meters == miles) // error, mismatched type
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
