package main

import "fmt"

func main() {

	// Interface
	/*
		Implementing an Interface:
		To implement the Speaker interface, a type just needs to implement the Speak() method.
	*/
	p := Person{Name: "abhiii71"}
	p.Speak()

	// Here Person has the method Speak(), so it implicitly implements the speaker interface. No explicit
	// declaration is needed.

	// POLYMORPHISM
	/*
		Polymorphism in Go is achieved through interfaces. It allows different types (structs) to implement the same interface and be used interchangeably,
		providing flexibility in the code. This is especially useful when you want to write functions that can operate on a variety of different types,
		as long as they share the same behavior (interface).
	*/
	a := Animal{Species: "Dog"}
	introduce(a)
	introduce(p)
	/*
		Polymorphism is demonstrated here: both Person and Animal implement the Speaker interface, but they provide their own implementations of the Speak() method.
		The function introduce(speaker Speaker) accepts any type that implements the Speaker interface, which in this case includes both Person and Animal.
		This allows you to write code that works with any type that satisfies the interface, without needing to know the specific type.
	*/

	// Empty Interface:
	// The empty interface is useful when you want to store any type of data, but you have to use type assertions or type switches to extract the underlying value.
	var x interface{}
	x = 42             // x can hold an int
	x = "Hello"        // x can hold a string
	x = []int{1, 2, 3} // x can hold a slice of ints
	fmt.Println("Interface: ", x)
}

// Interface
/*
An interface in Go defines a set of methods. Any type that implements these methods is considered to
satisfy the interface, allowing different types to share common behavior.
*/
/*
In the example above, we’ve defined an interface Speaker with a single method Speak().
Any type that has the method Speak() is said to implement the Speaker interface,
even though we don’t explicitly declare that the type implements the interface.
*/
type Speaker interface {
	Speak()
}

// Implementing an Interface:
// To implement the Speaker interface, a type just needs to implement the Speak() method.

// Person struct
type Person struct {
	Name string
}

// Implement to speak method for person
func (p Person) Speak() {
	fmt.Println("Hello my name is ", p.Name)
}

// Polymorphism
type Animal struct {
	Species string
}

// Implement speak to Animal
func (a Animal) Speak() {
	fmt.Println("Hello the species of Animal is ", a.Species)
}

func introduce(speaker Speaker) {
	speaker.Speak() // Call the speak method on any type that implements Speaker
}

// Key Points of Interfaces and Polymorphism:
/*
Interfaces in Go are implicit: A type doesn't need to explicitly declare that it implements an interface.
As long as it provides the methods defined by the interface, it implements it.
Interfaces provide flexibility: You can define functions that accept interfaces as arguments,
allowing them to operate on different types that implement the same interface.
Polymorphism allows you to call the same method (Speak() in our case) on different types (Person and Animal),
each of which can have a different implementation of that method.
*/

/*
// Type Assertion and Type Switch:
Type Assertion: Used to extract the concrete value from an interface.

var x interface{} = 42
v, ok := x.(int)  // v is 42, ok is true
Type Switch: Allows checking the actual type of an interface at runtime.

func printType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

printType(42)    // Integer: 42
printType("Go")   // String: Go
*/

/*
Interfaces allow you to define behavior without specifying how it is implemented.
Polymorphism allows different types to share the same interface and implement their own version of the interface methods.
Go’s interface system is implicit and flexible, making it easy to write generic functions that can operate on a variety of types.
The empty interface is a way to accept any type, and type assertions/type switches allow extracting the value or checking its concrete type.
*/
