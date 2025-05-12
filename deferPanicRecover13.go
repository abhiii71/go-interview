package main

import "fmt"

func main() {
	// 1. defer-delaying funtion execution until the end
	// Purpose: defer statement is used to delay the execution of a function until the surrounding function returns. This is commonly used to ensure that cleanup actions
	// such as closing files or releasing locks, occurs regardless of how the function exist(whether normally or due to a panic).
	// Behavior: when multiple defer statement are used, they are executed as LIFO(Last in First out) order, means last deferred function is executed first.
	fmt.Println("Starting the function")
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")

	// The deferred statement is delayed and will be executed when the function exists.
	// Key Point: The defer statement ensures that fmt.Println("This will be printed last")
	// is executed just before the function exits, even if an error occurs.

	// 2. Panic
	// 	result := divide2(1, 0) // this will cause panic
	// 	fmt.Println(result)     // so this will not  be going to execute

	// 3. Recover
	result := safeDrive(1, 0)
	fmt.Println("Result: ", result)
}

func divide2(a, b int) int {
	// 2. panic-Triggering a program failure
	// Purpose: The panic function is used to abort the execution of a program when something goes wrong
	// in an unexpected or unrecoverable way. When a panic occurs, it immediately stops the execution of the current function and begins unwinding the stack.
	// Common Usage: Typically used for fatal errors like invalid program states, broken invariants, or when continuing execution doesn't make sense.

	if b == 0 {
		panic("Division by zero!")
	}
	return a / b

	// When panic is called it triggers an immediate stop, and the program begins to unwind. If no recover is used, it will terminate the program
}

func safeDrive(a, b int) (result int) {
	// 3. Recover- Catching and Handling Panics

	// Purpose: The recover function is used to catch a panic and prevent the program from terminating. It can only be called within a defered function.
	// How it works: If a panic  is trigerred, Go begins to unwind the stack. If a recover function is called during the unwinding process(inside deferred function),
	// It catches the panic and returns the value passed to panic. This allows you to handle the panic is propagated through  the stack.
	// Key behavior: If recover is not called, the program will terminate after the  panic is propagated through  the stack
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
			result = 0 // Default result in case of panic
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b

	// Key Point: The recover function catches the panic, prevents the program from terminating, and provides a way to handle the error.
	// The function can then proceed, and the result is set to 0 as a default value.

}
