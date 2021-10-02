// This is a file header. It is separated from the package doc below with an
// empty line so it is ignored by Godoc. This area can be used for copyright
// notifications and/or internal descriptions of the file that you do not want
// to show.

// Package godoctricks is a tutorial that deals with tricks for making your
// godoc organized and neat. This is a compilation of tricks I've collected and
// couldn't find a comprehensive guide for.
//
// Notice that this doc is written in godoc itself as package documentation.
// The defined types are just for making the table of contents at the
// head of the page; they have no meanings as types.
//
// If you have any suggestion or comment, please feel free to open an issue on
// this tutorial's GitHub page!
//
// By Amit Lavon
package godoctricks

// You can see your godoc rendered as HTML by running a local godoc server.
// This is great for previewing your godoc before committing changes. To do
// that, Make sure your code is in GOPATH and run:
//  godoc -http ":8080"
// Go to http://localhost:8080/pkg and you should see your packages on the
// list.
//
// If you want the raw HTML, you can run:
//  godoc -url=/pkg/your_package > your_page.html
type HTML int

// Go code that you upload to public repositories on github appears
// automatically on the godoc website. Just like this tutorial! Just check in
// your code and watch as it appears. Use this page's URL as reference.
type Github int

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
// end with punctuation.
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
// to mock them (documentation-wise). Take this Mock_enums type for example,
// if you have a constant clause where all the constants are of the same type,
// it will be attached to that type's godoc. See below.
type Mock_enums int

const (
	A Mock_enums = 1
	B Mock_enums = 2
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
type Code_blocks int

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

// Functions that construct an instance of a type (or a pointer to it) are
// associated with the returned type.
func NewMethods() *Methods { return nil }

// You can mention bugs in the documentation. The syntax for that is like so:
//  // BUG(username): Some information.
//  // Some more information.
// This creates a section for bugs where each bug block is shown.
// You can use words other than "BUG", like "TODO". By default, only BUG notes
// are shown. If you run a godoc server locally, you can control that with
// the -notes argument. For example, -notes="BUG|TODO".
type Bugs int

// BUG(amit): This is an example bug.
// See the bugs section.

// You can run a local godoc server. This is helpful for previewing
// documentation, or for cases where you don't have a stable internet
// connection.
//
// First you need to get the godoc tool:
//  go get golang.org/x/tools/cmd/godoc
// Then simply running godoc will make it listen on localhost:6060.
// Run with -h to see how you can control the port number and other options.
type Local_server int
