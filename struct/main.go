package main

import "fmt"

/*
This example shows how to use struct types in Go.
*/

type BaseVehicle struct {
	brand string
	model string
}

// Simple struct
type Vehicle struct {
	brand      string
	model      string
	horsepower int
	year       int
}

func main() {
	initializeStruct()

	// use the composition
	createSpecificType()
}

func initializeStruct() {
	/*
		The following code shows how to initialize a struct
	*/
	// 1- Using the literal syntax
	fmt.Println(&Vehicle{brand: "Fiat",
		model:      "Punto",
		horsepower: 90,
		year:       2002})

	// 2- Using the the dot notation
	var v Vehicle
	v.brand = "Tesla"
	v.model = "Tesla1"
	v.horsepower = 120
	v.year = 2020
	fmt.Println(v)

	// 3- using literal in a particular order omitting the name of the fields
	fmt.Println(Vehicle{"Fiat", "Giulia", 140, 2018})

	// 4- using the new keyword that creates a pointer to the type initialing the fields
	// with the zero value of the specific type
	fmt.Println(new(Vehicle))
	// It is similar to
	fmt.Println(&Vehicle{})
}
