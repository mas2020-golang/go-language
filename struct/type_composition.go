package main

import "fmt"

/*
This example shows how to use the type composition to simulate inheritance.
*/

type BaseType struct {
	Name    string
	Surname string
	Age     int
}

func (b *BaseType) Greetings() {
	fmt.Println("Hello from BaseType!")
	b.Name = "test"
}

/*
Specific type contains the BaseType struct
*/
type SpecificType struct {
	*BaseType // using a pointer is better for performance reason
	Type      string
	Category  string
}

func createSpecificType() {
	// new SpecificType using the dot notation (ERROR if BaseType is a pointer)
	//var st SpecificType
	//st.Name = "My name"
	//st.Age = 12
	//st.Category = "Cat of spec type"
	//st.Type = "Type of spec type"
	//st.Surname = "Surname from Base Type"
	//fmt.Println(st)

	// new SpecificType using literal
	st1 := SpecificType{
		&BaseType{Name: "My name", Surname: "Surname from Base Type", Age: 12},
		"Type of spec type",
		"Cat of spec type",
	}
	// IMPORTANT: you can directly use the methods of the BaseType structure implicit included in the SpecificType
	st1.Greetings()
	fmt.Println("value of st1.BaseType is", st1.BaseType)
}
