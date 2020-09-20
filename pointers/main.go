package main

import "fmt"

/*
	This example shows how to use pointers in Go.
*/

// This is a simple struct that represents an artist
type Artist struct {
	Name, Genre string
	Songs       int
}

func main() {
	// this is an example to change the the value of a string using a pointer
	var ori string = "test"
	// pass the argument as a pointer to ori
	argumentsByRef(&ori)
	// ori has been changed by the function
	println(ori)

	/*
		This example changes a variable of a struct using a pointer. artist variable is a pointer.
		If we din't use a pointer the artist passed to the func would be a value (a copy) and not
		a reference.
	 */
	artist := &Artist{
		Name:  "test Name",
		Genre: "test",
		Songs: 12,
	}
	changeSongsNumber(100, artist)
	fmt.Println(artist)
}

/*
	This function as a parameter string that is a pointer to string. It means that
	changing the value to p will change the original value
*/
func argumentsByRef(p *string) {
	(*p) = "test2"
}

func changeSongsNumber(n int, artist *Artist){
	(*artist).Songs = n
}
