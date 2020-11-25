/*
Receiver of a struct
 */
package main

import "fmt"

// Simple struct
type VehicleNew struct {
	brand      string
	model      string
	horsepower int
	year       int
}

func (v *VehicleNew) Go() {
	fmt.Println("Started!")
}

func (v VehicleNew) GoSlow() {
	fmt.Println("Started!")
}

var color = struct {
	rgb int
	name string
}{
	100,
	"red",
}

func useVehicleNew(){
	// new object
	vn := VehicleNew{
		"Tesla", "TeslaXY", 125, 2020,
	}
	vn.Go() // Golang automatic converts to a pointer
	(&vn).Go() // explicit call using a Pointer
	var vnP *VehicleNew = &vn
	vnP.Go()

	vn2 := &VehicleNew{
		"Tesla", "TeslaXY", 125, 2020,
	}
	vn2.Go() // receiver is a pointer, argument vn2 is a pointer
	vn2.GoSlow() // argument is a pointer, Golang dereference the pointer automatically
	(*vn2).GoSlow() // explicit call dereferencing the pointer
}