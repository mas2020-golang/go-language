package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	// new instance of NewMyType
	m := NewMyType()
	fmt.Fprintf(&m, "Hello from %s", "MyNewType")

	// write examples
	writeWithWrite()
	writeWithFmt()

	// read examples
	readWithRead()
	readWithBufioReader(m)
	readAllBytes(m)

	// assignment test: this code will assign at the io.Reader (not addressable) interface the MyType object
	//var _ io.Reader = m -> compile error
	var _ io.Reader = &m
}

// Write directly using the Write method.
func writeWithWrite() {
	fmt.Println("---Using MyType.Write")
	m := NewMyType()
	m.Write([]byte{'a', 'b', 'c'})
	m.ShowContent()
	fmt.Println()
}

// Write using the Fprintf func.
func writeWithFmt() {
	fmt.Println("---Using fmt.Fprintf")
	m := NewMyType()
	fmt.Fprintf(&m, " - Hello from %s", "MyNewType")
	// Read data through the Stringer interface
	fmt.Println(m)
	fmt.Println()
}

// Use the Read interface to read 10 bytes at time
func readWithRead() {
	fmt.Println("---Using MyType.Read")
	m := NewMyType()
	fmt.Fprintf(&m, "Hello from %s", "MyNewType")
	//
	buffer := bytes.Buffer{}
	buf := make([]byte, 10)
	for true {
		n, err := m.Read(buf)
		if err == io.EOF {
			break
		}
		buffer.Write(buf[:n])
	}
	fmt.Printf("Read data from buffer: %s\n", buffer.Bytes())
	fmt.Println()
}

/*
Example of reading the MyType using the bufio.Reader.
*/
func readWithBufioReader(m MyType) {
	fmt.Println("---Using bufio.Reader")
	r := bufio.NewReader(&m)
	s, err := r.ReadString(' ')
	testError(err)
	fmt.Printf("string read through ReadString is %s\n", s)

	b, err := r.ReadBytes(10)
	testError(err)
	fmt.Printf("string read through ReadBytes is %s\n", b)
	fmt.Println()
}

/*
Read all the bytes.
*/
func readAllBytes(m MyType) {
	fmt.Println("---Using ioutil.ReadAll")
	data, err := ioutil.ReadAll(&m)
	testError(err)

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
	fmt.Println()
}

func useEmptyInterface() {
	var empty interface{}
	empty = 1
	empty = ""
	empty = true
	empty = map[int]int{1: 100}
	empty = MyType{}
	fmt.Println(empty)
}

func testError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
