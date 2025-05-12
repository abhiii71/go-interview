package main

import "fmt"

// 1. Struct
// A struct in Go is a composite data type (like a class without methods) that groups together fields of differnt types
type Person struct {
	Name string
	Age  int
}

func main() {
	// creating instance
	p := Person{Name: "Alice", Age: 23} // name fields
	_ = Person{"bob", 25}               // positional fields
	var r Person                        // zero value struct
	r.Name = "Charlie"
	r.Age = 22

	//Accessing and modifying fields
	fmt.Println(p.Name)
	p.Age = 31

	// 2. Pointers to structs
	// to avoid copying, structs are often used with pointers
	pp := &Person{Name: "Dan", Age: 33}
	fmt.Println("pp: ", pp.Name) // Automatically dereferenced
	pp.Age = 41

	// 3. struct with function
	pp.Greet()
	pp.HaveBirthday()
	fmt.Println("Update pp age: ", pp.Age)

	// 4. Struct Composition
	d := Dog{
		Animal: Animal{Name: "Fido"},
		Breed:  "Beagle",
	}

	fmt.Println(d.Name) // Access Animal.Name directly
	d.Speak()           // Access Animal method directly

	// Embedded structs export their fields and methods promoted to the outer struct

	// 5. Anonymous Structs
	// usesul for one-time use:

	pk := struct {
		Name string
		Age  int
	}{Name: "abhiii71", Age: 23}

	fmt.Println(pk)

	// 6. Comparing structs
	// structs can be compared only if all their fields are comparable:
	p1 := Person{"A", 71}
	p2 := Person{"A", 71}
	fmt.Println(p1 == p2) // true
}

// 3. Struct with functions(Methods)
// You can define methods on structs:

// receive by value (won't modify original )
func (p Person) Greet() {
	fmt.Println("Hello, my name is: ", p.Name)
}

// receive by reference(can modify orignal)
func (p *Person) HaveBirthday() {
	p.Age++
}

// 4. Struct Composition(Embedding)
// Go does composition over inheritance
type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("I am ", a.Name)
}

type Dog struct {
	Animal // embedded struct
	Breed  string
}

// 7. Struct Tags
// used for metadata(esp. for JSON, DB-mappings)
type Person1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 8. Zero values
// if you declare a struct without initializing fields, all fields get zero values:
// 0 for int
// "" for string
// nil for pointers, slices, etc.

// Why struct composition?
// promotes code reuse
// Enables methods promotion
// Preferred over traditional inheritance
// Provides a clean and flexible design in Go
