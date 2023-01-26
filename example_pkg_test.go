package godoctricks_test

import "fmt"

var message = "Hello"

// A whole file example is a file that ends in `_test.go` and contains exactly one
// example function, no test or benchmark functions, and at least one other
// package-level declaration.
func Example_wholeFilePackage() {
	fmt.Println(message)
	// Output: Hello
}
