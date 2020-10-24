package main

import "fmt"
import "sort"

func main() {
	fmt.Println("---Create maps")
	names := createMaps()
	fmt.Println("\n---Delete elements")
	deleteElements(names)
	//fmt.Println("\n---Add elements")
	addElements(names)
	fmt.Println("\n---Print elements")
	printValues(names)
	fmt.Println("\n---Sort a map")
	sortMap(names)
}

func createMaps() map[string]int{
	var testNil map[string]int // ZERO value is nil
	fmt.Printf("testNil == nil: %v, content is: %v\n", testNil == nil, testNil) // => true
	//testNil["one"] = 1 // panic

	// creation with make
	names := make(map[string]int)
	fmt.Printf("names == nil: %v, content is: %v\n", names == nil, names) // => false
	// maps from literals
	ages := map[string]int{
		"ag": 43,
		"sm": 29,
		"mg": 0,
	}
	fmt.Println("ages2 == nil:", ages == nil) // => false

	// another way to create an empty map
	ages2 := map[string]int{}
	fmt.Println("ages2 == nil:", ages2 == nil) // => false

	// assign values post creation
	names["ag"] = 43
	names["sm"] = 29
	// print a value
	fmt.Printf("names[\"ag\"]: %d\n", names["ag"])

	// check if a key exists:
	age, ok := names["test"]
	fmt.Printf("names test exists: %v and its value is %v", ok, age)

	return names
}

func deleteElements(m map[string]int) {
	delete(m, "ag")
	fmt.Printf("values of m are %v\n", m)
}

func addElements(m map[string]int) {
	m["test"] = 100
	m["zed"] = 90
	m["jhon"] = 53
}

func printValues(m map[string]int) {
	for name,age := range m {
		fmt.Printf("name: %s, age: %d\n", name, age)
	}
}

func sortMap(ages map[string]int)  {
	var names = make([]string, 0, len(ages)) // slice for the names
	// append all the names to the slice in a unordered way
	// we omit the second value in the for loop using only the key of ages and hot the value
	for key := range ages{
		names = append(names, key)
	}
	// sort slice of names
	sort.Strings(names)
	// read ordered names from slices and get the values of ages from the key name
	for _, name := range names{
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

