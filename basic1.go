package main

import "fmt"

func main() {
	/*
		// 1. Declaration Syntax
		// Go is a statically typed, meaning the type of a variable is known at compile time
		// explicit declaration
		var x int = 10

		// type inference
		var y = 10 // go infers y is int
		// short variable declaration (only inside function)
		z := 30

		// 2. Multiple Variable Declaration
		var a, b, c int = 1, 2, 3

		// or with type inference
		name, age := "Abhishek", 23

		// 3. Zero Values (Default Initialization)
		// uninitilized variables in Go have a default "zero value"
		// type		zero value
		// int		0
		// float64	0.0
		// bool		false
		// string	""(empty)
		// pointer	nil
		var xx int
		fmt.Println(x) // prints 0

		// 4. Constants
		// Declared using const, cannot be changed after assignment
		const pi = 3.14
		// constants must be assigned a value at compile time
		// can be typed  or untyped:
		const name string = "Go"
		const version = 1.20 // untype constant

		// data types
		// most same i mention here only few which have to focus more

		// string
		// immutable sequences of bytes
		var name string = "Abhishek"
		// strings can be concatenated:
		greeting := "Hello, " + name

		// Rune and byte
		// Byte
		// Alias for unit8 often used in slices
		var b byte = "A" // Ascii code of A = 65
		// rune
		// Alias for int32 represent unicode characters
		var r rune = '$'

		// Type Conversion
	*/
	// No implicit casting. Must use explicit conversion:
	var xy int = 10
	var yx float64 = float64(xy)
	fmt.Println(yx)
	// Type alias
	type Age int
	var myAge Age = 71
	fmt.Println(myAge)

	// Byte deep dive
	// Alias of uint8
	// typically used to work with binary data and Ascii characters
	var bb byte = 'A'      // Ascii of A is 65
	fmt.Println(bb)        // print 65
	fmt.Printf("%c\n", bb) // prints A

	//  In strings:
	s := "hello"
	b := s[0]             // b is byte
	fmt.Println(b)        // 104
	fmt.Printf("%c\n", b) // h

	// rune deep dive
	// Alias for int32
	// represents a unicode code point
	// useful for handling multibyte characters (like ₹, 你好, अ)
	var r rune = '$'
	fmt.Println(r) // uncode code point
	fmt.Printf("%c\n", r)
	// Multilingual character support

	// looping through a string
	ss := "#abc"
	// byte wise (may break multitype character)
	for i := 0; i < len(ss); i++ {
		fmt.Printf("ss[%d] = %c\n", i, ss[i]) // may print garbage for #
	}
	// rune-wise(Correct way)
	for i, r := range s {
		fmt.Printf("ss[%d] = %c\n", i, r)
	}
}
