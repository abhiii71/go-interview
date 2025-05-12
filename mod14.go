package main

func main() {
	//	1. go mod init <module_name>:
	/*
		This initializes a Go module in your project directory and creates a go.mod file.
		The go.mod file contains information about the module, including the name of the module and the list of dependencies it requires	.
		Example: go mod init mymodule
	*/

	//	2. go get <package>:
	/*
		This command is used to add a dependency to your module.
		When you run go get <package>, Go will:
		Fetch the specified package.
		Update the go.mod file with the new dependency.
		Add the package version to go.sum (a file that stores cryptographic hashes of dependencies for security).
		Example: go get github.com/gorilla/mux
		The go.mod file will be updated with the new dependency:
	*/

	//	3. go mod tidy:
	/*
		This command removes any dependencies that are no longer needed in your project.
		It also ensures that your go.mod and go.sum files are clean and contain only the required dependencies.
		It’s a good practice to run go mod tidy regularly to keep your dependencies up to date.
	*/

	//      4. go mod vendor:
	/*
		This command copies all of the module’s dependencies into a vendor/ directory within the project.
		It allows you to ensure that all dependencies are packaged along with your project, useful for
		situations where external sources might not be available or for specific build environments.
	*/

	//	Managing Dependencies with Go Modules
	/*
							Adding Dependencies:/ Updating Dependencies:
							You can add dependencies using go get <package>. For example, if you want to use the mux router package from
							github.com/gorilla/mux, you can run:

						//	To update all dependencies, you can use:
							go get -u

			// Checking for Dependency Issues:
				The command go mod verify can be used to check the integrity of the dependencies in your module:
				go mod verify

		// List dependencies
		go list -m all


	*/
}
