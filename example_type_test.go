package godoctricks_test

import "fmt"

var messageExample = "Hello"

// A whole file example can also be attached to a given type (Examples here)
// with appropriate naming.
//
// A whole file example is a file that ends in `_test.go` and contains exactly one
// example function, no test or benchmark functions, and at least one other
// package-level declaration.
func ExampleExamples_wholeFileType() {
	fmt.Println(messageExample)
}
