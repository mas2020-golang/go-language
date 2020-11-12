package main

import (
	"bytes"
	"fmt"
	"strconv"
)
import "unicode/utf8"

func main() {
	//slice()
	//fmt.Println("---Rune")
	//decodeRunes()
	//fmt.Println("---Unicode Escapes")
	//unicodeEscapes()
	//fmt.Println("---Get Runes")
	//getRunes()
	//fmt.Println(comma("123456789"))
	conversions()
	fmt.Printf("int into string: %v", useByteBuffer([]int{65, 66, 67}))
}

func slice() {
	s := "Hello Golang Aware!"
	fmt.Println(s[:])   // -> Hello Golang Aware!
	fmt.Println(s[:5])  // -> Hello
	fmt.Println(s[2:5]) // -> llo
}

func decodeRunes() {
	s := "◊Hello, ◊Ç∞"
	fmt.Println("bytes length:", len(s))                      // "15"
	fmt.Println("rune in string:", utf8.RuneCountInString(s)) // "11"

	// Decoding rune: scan of a string reading the character based on the position due to the size
	// of the previous char
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Println()
	// decodes implicitly
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

}

func unicodeEscapes() {
	// unicode escapes allow to write a unicode characters using its numeric code point instead of string literal
	// for example it is useful where a character is not easy to write
	// error character
	fmt.Printf("%c\t%d\n", '\uFFFD', '\uFFFD')                             // -> �, unicode escape into a rune, code point 65533
	fmt.Printf("%c\t%d\n", '\u4e16', '\u4e16')                             // -> 世, unicode escape into a rune, code point 19990
	fmt.Printf("%s\tlen: %d bytes\n", "\xe4\xb8\x96", len("\xe4\xb8\x96")) // -> 世 bytes: 3, hexadecimal escape, each \xx represents 1 byte
	fmt.Printf("%s\tlen: %d bytes\n", "\u4e16", len("\u4e16"))             // -> 世 bytes: 3, unicode escape into a string
}

func getRunes() {
	s := "世±—·"
	fmt.Printf("% x\n", s) // --> hexadecimal format where each bytes is separated by space
	r := []rune(s)
	fmt.Printf("%x\n", r) // --> hexadecimal format
	fmt.Printf("%v\n", r) // --> uint32 format
	var n uint32 = uint32(r[0])
	fmt.Printf("uint32 code point for r[0] is: %d", n)
}

// comma inserts commas in a non-negative decimal integer string
// using recursion (example extracted) from the book 'Go Programming Language"
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Example of some way to convert from and into strings
func conversions() {
	fmt.Println("---Conversions")
	s := "This is my string世"
	// convert into a slice of bytes (opposite to string cab be changed)
	b := []byte(s)
	fmt.Printf("bytes len %d\t%v\n", len(b), b) // --> 20, the last char occupies 3 bytes
	fmt.Printf("hex digits % x\n", s)           // --> 20, the last char occupies 3 bytes
	// convert from bytes to string
	s2 := string(b)
	fmt.Printf("original string:\t%s\nconverted string:\t%s\n", s, s2)

	// Convert from int to string
	n := 123
	fmt.Printf("str with Sprintf() is %s\n", fmt.Sprintf("%d", n))        // -> 123, string rep of an int
	fmt.Printf("str binary with Sprintf() is %s\n", fmt.Sprintf("%b", n)) // -> 1111011, string rep of an int converted to binary
	fmt.Printf("str with string() is %s\n", string(n))                    // -> {, because convert from int to UTF-8 representation
	fmt.Printf("str with strconv.Itoa() is %s\n", strconv.Itoa(n))        // // -> "123"

	// Convert from string to int
	n, _ = strconv.Atoi("123")              // -> conversion using strconv
	m, _ := strconv.ParseInt("123", 10, 16) // -> conversion using ParseInt, return i64 into n
	fmt.Println("n,m", n, m)
}

// Example on how to use ByteBuffer
func useByteBuffer(v []int) string {
	var buf *bytes.Buffer = &bytes.Buffer{}
	buf.WriteByte('[')
	for i, v := range v {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
