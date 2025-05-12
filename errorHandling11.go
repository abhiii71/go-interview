package main

import (
	"errors"
	"fmt"
)

func main() {
	// ERROR HANDLING

	//1. ther error type
	// in go, error is a built in interface
	// Any type that implements the Error() string method satisfies the error interface
	type error interface {
		Error() string
	}

	// 2. Creating Errors
	// Using errors.New()
	// From the errors package:
	// import "errors"
	// This returns a simple error with a static message
	err := errors.New("Something went wrong")
	fmt.Println(err)

	// Using fmt.Errorf()
	// Used for formatted error message:
	// import "fmt"
	// fmt.Errorf() is powerful when you need dynamic context in your errors
	name := "file.txt"
	err = fmt.Errorf("failed to open file %s", name)

	// 3. Function that returns(value, error)
	fmt.Println("Function that returns(value, error): ")
	// handling it
	// Always check if err != nil before using the result
	result, err := dividee(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

}

// ERROR HANDLING
// 3. Functions that Return (value, error)
// Go favors returning an error as the last return value of a function:
func dividee(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil

}

// 4. Custom Error Types(Bonus)
// you can define your own error struct:
// This  gives you more power to distinguish  error types
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Msg)
}

func fail() error {
	return &MyError{Code: 500, Msg: "Internal Server Error"}
}

// 5. Best practices
/*
Always return error as the last return value
Always check for nil before using the result
Prefer fmt.Errorf() when you want context
Use custom error types if you want to match error values with errors.As() or error.Is()(advanced usage)
*/
