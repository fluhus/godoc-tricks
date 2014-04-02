// This tutorial deals with tricks for making your godoc organized and neat.
// This is a compilation of tricks I've collected and couldn't find a
// comprehensive guide for.
//
// Notice that this doc is written in godoc itself as a package's documentation.
// The defined types are just for making the table of contents at the
// head of the page. They have no meanings as types.
//
// Author: Amit Lavon
package godoctricks

// You can output an HTML page if you want a fancier, JavaDoc-like output.
// For that, make sure GOPATH is set correctly and run the command:
//  godoc -url=/pkg/your_package > your_page.html
// Now if you don't have a godoc server, you'll get an ugly looking page,
// since there is no CSS and no script files. Link your HTML
// to some godoc website's CSS and JS files. I add http://golang.org/ as prefix
// to the CSS file address (found at head section) and the script files (found
// at the bottom of the page's source).
type HTML int

// To start a new paragraph, add an empty line in the comment between the 2
// paragraphs.
//
// For example:
//  // Paragraph 1.
//  // Still paragraph 1.
//  //
//  // Paragraph 2.
//  // Still Paragraph 2.
// will yield:
//
// Paragraph 1.
// Still paragraph 1.
//
// Paragraph 2.
// Still Paragraph 2.
type Paragraphs int

// You can make titles in your godoc. A title is a line that is separated from
// its following line by an empty line, begins with a capital letter and doesn't
// end with panctuation.
//
// For example, the code:
//  // Sentence 1
//  //
//  // Sentence 2
// will yield:
//
// Sentence 1
//
// Sentence 2
//
// While this code:
//  // Sentence 1.
//  //
//  // Sentence 2.
// will yield:
//
// Sentence 1.
//
// Sentence 2.
//
// See documentation here: http://golang.org/pkg/go/doc/#ToHTML
type Titles int

// While there are no built in enums in go, you can use types and constants
// to create them. Take this Enums type for example, if you have a constant
// clause where all the constants are of the same type, it will be attached
// to that type's godoc. See below.
type Enums int

const (
	A Enums = 1
	B Enums = 2
)

// You can place usage examples in your godoc.
//
// Examples should be placed in a file with a _test suffix. For example, the
// examples in this guide are in a file called doc_test.go .
//
// The example functions should be called Example() for
// package examples, ExampleTypename() for a specific type or
// ExampleFuncname() for a specific function. For multiple examples
// for the same entity (like same function), you can add a suffix like
// ExampleFoo_suffix1, ExampleFoo_suffix2.
//
// You can document an example's output, by adding an output comment at its end.
// The output comment must begin with "Output:", as shown below:
//  func ExampleExamples_output() {
//      fmt.Println("Hello")
//      // Output: Hello
//  }
//
// Notice that the tricks brought here (titles, code blocks, links etc.) don't work
// in example documentation.
//
// For full documentation of examples, see:
// http://golang.org/pkg/testing/
type Examples int

// You can embed blocks of code in your godoc, such as this:
//  fmt.Println("Hello")
// To do that, simply add an extra indent to your comment's text.
//
// For example, the code of the first lines of this section looks like this:
//  // You can embed blocks of code in your godoc, such as this:
//  //  fmt.Println("Hello")
//  // To do that, simply add an extra indent to your comment's text.
type CodeBlocks int

// Web addresses will automatically generate actual links in the HTML output,
// like this: http://www.golang.org
type Links int

// Methods are functions with receivers. Godoc associates methods with their
// receivers and attaches their documentation. See below.
type Methods int

// Methods are attached to their receiver type in the godoc, regardless of
// their physical location in the code.
func (Methods) Foo() {}

// Pointer receivers are also associated in the same way.
func (*Methods) Foo2() {}

// Functions that return a type (or a pointer to it) are also associated
// with it.
func NewMethods() *Methods {}

