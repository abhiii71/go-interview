package utils

import "fmt"

// Export function(capitalized)
func Add(a, b int) int {
	return a + b
}

// Unexported function (not accessible from outside)
func log(msg string) {
	fmt.Println("Log: ", msg)
}
