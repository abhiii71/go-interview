package main

import (
	f "fmt"
	// you can use alias import as well
	"myapp/utils"
)

func main() {
	result := utils.Add(6, 1)
	f.Println("Sum: ", result)
	// Creating Your Own Packages
	// Function and variable names must be capitalized to be exported(accessible from other packages)
	// Directory name != Package name (you can name the package differently inside the .go file)
	// import path is based on module path + directory structure

	// Go Modules(organizing code in real projects)
	// 1. initialize module:
	// go mod init myapp
	// 2. Now Go knows your base path is myapp, and you can import internal packages using:
	// import "myapp/utils"
	// 3. Build/run the project:
	// go run main.go

	// Advanced Notes for interview
	// init() functions: Automatically invoked before main()or when a package is imported
	// Relative imports are not allowed -Go only uses module path or absolute import paths.
	// Tools like go install , go build, go test behaves differently when working with packages

}
