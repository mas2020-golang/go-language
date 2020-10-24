package main

import "fmt"

func main() {
	// create slices example
	createSlices()
	// copy slices example
	copySlices()
	// shows how to append elements
	appendElement()
	// show how to delete an element
	deleteElement(1)
	deleteElementWithOrder(1)
	checkNilValues()
	// examples of slicing
	slicesOperations()

	fmt.Println("\n---Reverse example")
	// int slice and reverse (Go passes a copy of a pointer of the slice, so the func can change the values)
	i := []int{1, 2, 3, 4, 5, 6}
	reverse(i)
	fmt.Println(i)
}

func slicesOperations(){
	fmt.Println("\n---Some slice operations")
	// this implicit initialization specifies the number of index associate to a value. In this case 0 and 2 are
	// initialized has an empty string (ZERO value for the string is "")
	months := []string{1: "January", 3: "March", "April", "May", "June", "July"}
	printValues(months)
	fmt.Printf("type is %T, len is %v, cap is a %v\n", months, len(months), cap(months)) // => 4, 4
	fmt.Printf("type is %T, len is %v, cap is a %v\n",months,  len(months), cap(months)) // => 4, 4

	// an example of the slice operator that creates a new slice
	firstMonths := months[0:3] // -> 0,1,2 indexes
	printValues(firstMonths)
	// other slice examples
	allMonths := months[:]
	// from beginning to end
	allMonthsTwo := months[0:]
	// or
	allMonthsThree := months[0:8]
	printValues(allMonths)
	printValues(allMonthsTwo)
	printValues(allMonthsThree)

	// to extend a slice always inside the capacity
	fourMonths := firstMonths[:4]
	printValues(fourMonths)
}

func checkNilValues(){
	fmt.Println("\n---Check nil values")

	// int slice and reverse (Go passes a copy of a pointer of the slice, so the func can change the values)
	i := []int{1,2,3,4,5,6}
	reverse(i)
	fmt.Println(i)

	// nil values
	var s []int
	fmt.Printf("len(s) is %v, s == nil; %v\n", len(s), s == nil) // len(s) is 0, len == nil; true
	s = nil
	fmt.Printf("len(s) is %v, s == nil; %v\n", len(s), s == nil) // len(s) is 0, len == nil; true
	s = []int(nil)
	fmt.Printf("len(s) is %v, s == nil; %v\n", len(s), s == nil) // len(s) is 0, len == nil; true
	s = []int{}
	fmt.Printf("len(s) is %v, s == nil; %v\n", len(s), s == nil) // len(s) is 0, len == nil; false
}

/*
Some ways to create a slice.
*/
func createSlices() {
	fmt.Println("---Create slices")
	var s []int               // nil slice
	var s1 = []int{}          // empty slice with 0 elemes (not nil)
	s2 := []int{1, 2, 3}      // slice literal
	s3 := make([]int, 10, 20) // slice of 10 len and 20 cap with 10 elems set to int ZERO value
	s4 := new([20]int)[:10]   // same as s3
	fmt.Println(s, s1, s2, s3, s4)
}

/*
Some ways to create a slice.
*/
func copySlices() {
	fmt.Println("\n---Copy two slices")
	var s3 = make([]int, 2)
	s2 := []int{1, 2, 3, 5, 6} // slice literal
	n := copy(s3, s2)
	fmt.Printf("number of elements copied is %d, value of dest is: %v\n", n, s3)

	// creation with make (create a variable with all the elements of ZERO types each one)
	b := make([]int, 10, 20)
	fmt.Println(b)

	// shows how to append elements
	appendElement()
}

func printValues(a []string) {
	fmt.Println("=> Values of the param a are:")
	for i, v := range a {
		fmt.Printf("i: %d, value: %s\n", i, v)
	}
	fmt.Println()
}

/*
Delete an element changing the order.
*/
func deleteElement(pos int){
	fmt.Println("\n---Delete element")
	s := []int{10,20,30,50,100,90}
	fmt.Println("before deletion s slice is", s)

	// Remove the element at index i from a.
	s[pos] = s[len(s)-1] // last element overrides the pos passed
	s = s[:len(s)-1]   // remove last element
	fmt.Printf("after deletion s slice is %v (len: %d, cap: %d)\n", s, len(s), cap(s))
}

/*
Delete an element maintaining the order.
*/
func deleteElementWithOrder(pos int){
	fmt.Println("\n---Delete element maintaining the order (slower than deleteElement)")
	s := []int{10,20,30,50,100,90}
	fmt.Println("before deletion s slice is", s)
	copy(s[pos:], s[pos+1:]) // shift left of one pos
	s[len(s)-1] = 0 // ZERO value on the last element
	s = s[:len(s)-1]   // remove last element
	fmt.Printf("after deletion s slice is %v (len: %d, cap: %d)\n", s, len(s), cap(s))
}

func reverse(a []int){
	for i,j := 0, len(a) - 1; i < j; i,j = i+1, j-1{
		a[i], a[j] = a[j], a [i]
	}
}

func appendElement(){
	// Creation of a slice of 5 elements
	i := make([]int, 5, 10)
	fmt.Printf("type is %T, len is %v, cap is a %v\n",i,  len(i), cap(i)) // => 4, 4
	fmt.Println(i) // => len == 5
	i = append(i, 100)
	fmt.Println(i) // len == 6, cap == 10 => [0 0 0 0 0 100]
}
