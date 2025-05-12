package main

import "fmt"

func main() {

	// ARRAYS

	// arrays are fixed size, same-type collections.
	// Once created, you cannot change their length
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// can ommit the array values:
	var arrr [3]int
	fmt.Println(arrr)

	// shorthand declaration
	ar := [3]int{1, 2, 3}
	fmt.Println(ar)

	// Compiler infers size:
	ar1 := [...]int{1, 2, 3}
	fmt.Println(ar1)

	// Access and iteration
	for i, v := range ar1 {
		fmt.Println("index: ", i, " value: ", v)
	}

	// LIMITATION
	// Cannot resize after declaration
	// Passed by value to functions(copy is made)

	// SLICE

	// Dynamic, flexible view over an array
	// Internally: pointer to array+lenght+capacity
	// Most used collection type in Go

	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	// create using make
	_ = make([]int, 3) // lenght = 3, capacity = 3

	_ = make([]int, 3, 5) // length = 3, capacity = 5

	//Created from array
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice := arr1[1:4] // includes index 1, excludes 4 => [2, 3, 4]
	fmt.Println("slice: ", slice)

	// Slice operations
	// 1. Append
	s1 := []int{1, 2}
	s1 = append(s1, 3)                 // append the array with [1, 2, 3]
	s1 = append(s1, []int{4, 5, 6}...) // variadic append

	// Returns a new slice (can point to new array if capacity is exceeded)

	// 2. Lenght and Capacity
	s2 := []int{1, 2, 3}
	fmt.Println(len(s2))                 //3
	fmt.Println("Cap before: ", cap(s2)) // 3 (depends on underlying array)

	s2 = append(s2, 4)
	fmt.Println("Cap  after append: ", cap(s2)) // Might increase (usually doubles)

	// slicing-a slice
	a := []int{1, 2, 4, 5, 71}
	b := a[2:5]
	fmt.Println("Slice: ", b)

	// No copy is made both a and b share the same underlying array.
	// use copy (dst, src) to copy slices
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(src, dst)

	// 4. Important Slice internals
	slicee := make([]int, 3, 5)
	// lenght : 3(number of accessible elements)
	// Capacity: 5(underlying array can hold up to 5 before reallocating)

	// appending beyond capacity creates a new underlying array:
	slicee = append(slice, 4, 5)             // okay
	fmt.Println("Cap before: ", cap(slicee)) // okay
	slice = append(slice, 6)                 // trigger a new array allocation
	fmt.Println("Cap after: ", cap(slicee))

	// Multidimentional Arrays and Slices
	// array
	var ar12 [2][3]int
	ar12[0][1] = 5

	//slices:
	matrix := [][]int{
		{1, 3, 5},
		{7, 9, 11},
	}
	fmt.Println("Slice Matrix: ", matrix)
	// More flexible than multidimensinal arrays
	// rows can be of different lengths

	// MAP

	// 1. Declartion and Initialization
	// Syntax:
	// a). Using a map literal:
	m := map[int]string{1: "abhiii71"}
	//store  things in key and value
	// type of key is integer and value is type of string
	fmt.Println("Map: ", m)

	// b). Using make() (empty map):
	m1 := make(map[int]string)
	// this creates an empty map with keys of type string and values of type int. You can add key-value pairs later.
	fmt.Println("Empty map: ", m1)

	// c). Using make() with intitla capacity:

	m2 := make(map[string]int, 10) // initial capacity of 10 elements
	// This initializes an empty map but with a specified initial capacity(which helps optimize performance by reducing reallocations during insertions.)
	fmt.Println("Map2: ", m2)

	// 2. Get and Set Values
	// Set(Add or Update Key-Value Pair):
	m1[1] = "one"
	m1[2] = "two"
	// Adds a new key-value pair or updates the value if the key already exists.

	//Get(Access a value by key)
	// if key doesn't exist, it returns the zero value for the value type (according to type like int stirng )
	value := m1[1]
	fmt.Println("Map1 operation: ", value) // "one"

	// Check if key exists:
	// ok is true if key exists in the map , and false otherwise
	value, ok := m1[1]
	if ok {
		fmt.Println("Printing the value: ", value)
	} else {
		fmt.Println("Key not found")
	}

	// 3. Delete Values
	// you can delete a key-value pair from a map using the delete() function:
	// This removes the key 1 from the map m1
	// if the key doesn't exists, nothing happens (i.e, no error)
	fmt.Println("Map m1 before delete: ", m1)
	delete(m1, 1)
	fmt.Println("Map m1 after delete: ", m1)

	// 4. Looping Over Maps(Using for range)
	// Basic Iteration:
	// Order is not guaranteed when looping through a map in Go. The iteration order is random each time.
	m3 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Iterating over the map: ")
	for key, value := range m3 {
		fmt.Println(key, value)
	}

	// Iterating over only keys or values:
	// looping over keys :
	for key := range m3 {
		fmt.Println("Keys: ", key)
	}

	// over values only:
	for _, values := range m3 {
		fmt.Println("Values: ", values)
	}

	// 5. Important Consideration
	// Zero value of Maps:
	// The zero value for a map is nil. A nil map behaves like an empty map when reading but causes a runtime panic if you try to write it.
	var m4 map[string]int
	fmt.Println(m4 == nil) // true
	// m4["key"] = 10         // this would panic because the map is nil
	fmt.Println(m4)

	// Map lenght
	// same function as array
	fmt.Println("length of m4: ", len(m4))

}
