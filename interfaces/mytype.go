package main

import (
	"bytes"
	"fmt"
	"io"
)

type MyType struct {
	buffer bytes.Buffer
	i int64
}

// "Constructor" for our type
func NewMyType() MyType{
	//m := MyType{
	//	buffer: bytes.Buffer{},
	//}
	//m.buffer.Grow(1000)
	return MyType{}
}

// Method that implements io.Writer interface
func (t *MyType) Write(p []byte) (n int, err error){
	return t.buffer.Write(p)
}

// Method that implements io.Reader interface
func (t *MyType) Read(p []byte) (n int, err error){
	// Check if we are at the end
	if t.i >= int64(t.buffer.Len()) {
		return 0, io.EOF
	}
	n = copy(p, t.buffer.Bytes()[t.i:])
	t.i += int64(n)
	return
}

// Show the content of our type as string
func (t *MyType) ShowContent() {
	fmt.Println(t.buffer.String())
}

// Show the content of our type as string using the Stringer interface
func (t MyType) String() string {
	return t.buffer.String()
}
