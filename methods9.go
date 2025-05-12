package main

import "fmt"

// Methods
/*
In Go, methods are functions that are associated with a specific type, typically a struct.
These methods allow you to define behavior (functions) that can be executed on instances of a struct.
The syntax for defining methods in Go is slightly different from that of regular functions because they have a "receiver."

Key Points:
Receiver: A method has a special parameter called the "receiver," which binds the method to a particular type (like a struct).
It behaves similarly to this or self in other languages.
Method Signature: The method signature includes the receiver, which is placed before the method name.
Value vs Pointer Receiver: You can define methods with either value receivers or pointer receivers,
depending on whether you want the method to modify the struct itself or just work with a copy of it.
*/

// Define a struct named Person
type Personn struct {
	Name string
	Age  int
}

// method with value receiver
func (p Personn) Greet() {
	fmt.Println("Hello", p.Name)
}

// Method receiver a pointer receiver
func (p *Personn) HaveBirthDay() {
	p.Age++
	fmt.Println("Happy Birthday! ", p.Name, " is now ", p.Age)
}

/*
Value Receiver: The method works on a copy of the struct. Any changes to fields inside the method will not affect the original struct outside the method.

Pointer Receiver: The method works on the actual struct (via pointer), and it can modify the original struct's fields.
*/
func main() {
	// create an instance of Person
	p := Personn{"abhiii71", 70}

	// call the method Greet()on the instance
	p.Greet()
	// Call the method HaveBirthday() on the instance
	p.HaveBirthDay()

	// Call the method again after incrementing the age
	p.HaveBirthDay()
}

// Example with Interface:
/*
// Define a struct
type Car struct {
    Model string
}

// Define an interface with a method
type Vehicle interface {
    Start()
}

// Method attached to Car type
func (c Car) Start() {
    fmt.Println("The", c.Model, "is starting!")
}

func main() {
    // Create an instance of Car
    myCar := Car{Model: "Tesla"}

    // Use the Car type as a Vehicle
    var v Vehicle = myCar
    v.Start()  // Calls the Start method of Car
}
*/
