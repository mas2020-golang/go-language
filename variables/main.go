package main

import "fmt"

/*
	This example shows how to create variables and constants in Go.
 */

func main(){
	variables()
	constants()
}

func variables(){
	fmt.Println("==> Variables section:")
	// implicit assignment (declaration and assignments are done at the same time
	name, location := "Prince Oberyn", "Dorne"
	fmt.Println("variables are:", name, location)

	// explicit assignment
	var (
		name1     string
		age1      int
	)
	name1, age1 = "test", 1
	fmt.Println("variables are:", name1, age1)

	// declaration of multiple variables
	var a,b,c int
	println("variables a,b,c, are:", a,b,c)

	// Variable can also be declared one by one
	var one string
	var second int
	var third bool = false
	one, second = "hello", 100
	fmt.Println("variables are:", one, second, third)
}

func constants(){
	/*
		Constants can only be character, string, boolean, or numeric values.
		Cannot be declared using the := syntax
	*/
	fmt.Println("==> Constants section:")
	const Pi = 3.14
	const (
		StatusOK                   = 200
		StatusCreated              = 201
		Big   = 1 << 62
	)
	fmt.Printf("Constants values are: Pi = %f, StatusOK = %d, StatusCreated = %d" +
		", Big = %d", Pi, StatusOK, StatusCreated, Big)
}


