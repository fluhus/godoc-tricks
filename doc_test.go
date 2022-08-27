package godoctricks

import (
	"fmt"
)

// This function is named Example(), which is a package level test
// function.
func Example() {
	fmt.Println("Hello")
}

// This function is named Example_other(), which is a package level test
// function and has the label "other".
func Example_other() {
	fmt.Println("Hello")
}

// This function is named ExampleExamples(), this way godoc knows to associate
// it with the Examples type.
func ExampleExamples() {
	fmt.Println("Hello")
}

// This function is named ExampleExamples_other(), it is associated with
// Examples type and has the label "Other".
func ExampleExamples_other() {
	fmt.Println("Hello")
}

// This is how godoc parsed ExampleExamples_output() that was shown above.
func ExampleExamples_output() {
	fmt.Println("Hello")
	// Output: Hello
}
