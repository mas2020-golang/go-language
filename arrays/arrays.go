package main

import "fmt"

func main() {
	var a [5]string // declaration of an array of 5 elements (init at ZERO values)
	// cycle to get the values
	printValues(a)
	// set first and last value
	a[0] = "first"
	a[len(a) - 1] = "last"
	printValues(a)

	// initialization using literals
	var b [5]string = [5]string{"a", "b", "c", "d", "e"}
//	b[5] = "test" // compile error => out of bounds
	printValues(b)

	// initialization using literals (without var) taking the length from the
	// number of elements
	c := [...]string{"a", "b", "c", "d", "e"}
	printValues(c)

	// it is possible to compare only two arrays of the same type and the same length
	// two arrays are equals if the elements have the same values
	fmt.Println("b == c", b == c)

	i1 := [3]int{1,2,3}
	i2 := [3]int{1,2,4}
	i3 := [4]int{1,2,4,5}
	fmt.Println("i1 == i2", i1 == i2)
	fmt.Println("i3:", i3)
//	fmt.Println("i1 == i3", i1 == i3) // compile time error


}

func printValues(a [5]string) {
	fmt.Println("=> Values of param [5]string are:")
	for i, v := range a {
		fmt.Printf("i: %d, value: %s\n", i, v)
	}
	fmt.Println()
}
