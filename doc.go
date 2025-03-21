// This is a file header. It is separated from the package doc below with an
// empty line so it is ignored by Godoc. This area can be used for copyright
// notifications and/or internal descriptions of the file that you do not want
// to show.

// Package godoctricks is a tutorial on the features of GoDoc.
// It is meant to help eveyone make the most out of this feature of Go.
//
// Notice that this doc is written in godoc itself as package documentation.
// The defined types are just for making the table of contents at the
// head of the page; they have no meanings as actual types.
//
// If you have any suggestion or comment, please feel free to [open an issue].
//
// By Amit Lavon.
//
// [open an issue]: https://github.com/fluhus/godoc-tricks/issues
package godoctricks

import "go/doc"

// You can run a local godoc server. This is helpful for previewing
// documentation, or for cases where you don't have a stable internet
// connection.
//
//	go run golang.org/x/tools/cmd/godoc@latest
//
// godoc will listen on http://localhost:6060.
// Run with -h to see how you can control the port number and other options.
type LocalServer int

// Go code that you upload to public repositories on github appears
// automatically on the godoc website. Just like this tutorial! Just check in
// your code and watch as it appears. Use this page's URL as reference.
//
// The godoc page gets updated once you make a new release.
type Github int

// To start a new paragraph, add an empty line in the comment between the 2
// paragraphs.
//
// For example:
//
//	// Paragraph 1.
//	// Still paragraph 1.
//	//
//	// Paragraph 2.
//	// Still Paragraph 2.
//
// Results in:
//
// Paragraph 1.
// Still paragraph 1.
//
// Paragraph 2.
// Still Paragraph 2.
type Paragraphs int

// You can make headings in your godoc.
// An explicit heading line starts with a '#' and is separated from the previous
// and next paragraphs with empty comment lines.
// An implicit heading is the same without the '#' (`go fmt` will automatically add a leading #).
// It needs to begin with a capital letter and not end with punctuation.
//
// See the full documentation here: https://go.dev/doc/comment#headings
//
// For example, this code:
//
//	// # Explicit Heading
//	//
//	// Yada yada.
//	//
//	// Implicit Heading
//	//
//	// Yolo yolo.
//
// Results in:
//
// # Explicit Heading
//
// Yada yada.
//
// Implicit Heading
//
// Yolo yolo.
type Headings int

// While there are no built in enums in go, you can use types and constants
// to make enum lookalikes (documentation-wise). Take this Enums type for
// example; if you have a const/var clause where all the values are of the same
// type, it will be attached to that type's godoc. See below.
type Enums int

const (
	A Enums = iota
	B
)

// You can place usage examples in your godoc.
//
// Examples should be placed in a file with a _test suffix. For example, the
// examples in this guide are in a file called doc_test.go .
//
// The example functions should be called Example() for
// package examples, ExampleTypename() for a specific type or
// ExampleFuncname() for a specific function, ExampleTypename_Methodname() for
// a method on a specific type.
//
// For multiple examples for the same entity (like same function), you can add
// a suffix starting with a lowercase like ExampleFoo_suffix1, ExampleFoo_suffix2.
//
// You can document an example's output, by adding an output comment at its end.
// The output comment must begin with "Output:", as shown below:
//
//	func ExampleExamples_output() {
//	    fmt.Println("Hello")
//	    // Output: Hello
//	}
//
// Notice that the tricks brought here (headings, code blocks, links etc.)
// don't work in example documentation.
//
// For full documentation of examples, see:
// http://golang.org/pkg/testing/
type Examples int

// NewExamples is a func with an associated example. See [Examples] regarding the details.
func NewExamples() Examples {
	return 0
}

// ExampleMethod is a method with an associated example. See [Examples] regarding the details.
func (e Examples) ExampleMethod() int {
	return int(e)
}

// FuncExampleWithoutType is a func with an associated example (but not associated type).
// See [Examples] regarding the details.
func FuncExampleWithoutType() int {
	return 0
}

// You can embed blocks of code in your godoc, such as this:
//
//	fmt.Println("Hello")
//
// To do that, simply add an extra indent to your comment's text.
//
// For example, the code of the first lines of this section looks like this:
//
//	// You can embed blocks of code in your godoc, such as this:
//	//  fmt.Println("Hello")
//	// To do that, simply add an extra indent to your comment's text.
type CodeBlocks int

// There are several ways to embed links in godoc.
//
// # Regular URLs
//
// Web addresses will automatically generate links in the HTML output,
// like this: http://www.golang.org
//
// # Documentation Links
//
// You can link to identifiers in the documentation with square brackets.
//   - [Name] or [Name.Name] for a member of the current package
//   - [pkg], [pkg.Name] or [pks.Name.Name] for a member of an imported package
//   - [import/path/pkg], [import/path/pkg.Name] or [import/path/pkg.Name.Name]
//     for a member of a non-imported package
//
// # Markdown-Style Links
//
// Lines of the form "[something]: URL" turn the occurrences of [something]
// into links to that URL with the text "something".
// Lines of that form need to be separated from regular text with empty doc
// lines.
//
// For example:
//
//	You can search for Go [online].
//	See our [Links] section or [Methods.Foo].
//	Go to the imported [doc] package or the [doc.Note] struct.
//	Visit us under [github.com/fluhus/godoc-tricks.Links]
//
//	[online]: https://duckduckgo.com/?q=golang
//
// Results in:
//
// You can search for Go [online].
// See our [Links] section or [Methods.Foo].
// Go to the imported [doc] package or the [doc.Note] struct.
// Visit us under [github.com/fluhus/godoc-tricks.Links].
//
// [online]: https://duckduckgo.com/?q=golang
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

// Struct will be shown with their exported fields. Unexported fields
// will be hidden under a comment "contains filtered or unexported fields".
// For instance the following struct
//
//	type Struct struct {
//		// ExportedField can have a comment on the line above.
//		ExportedField       int
//		OtherExportedField  Links // or on the same line
//		FieldWithoutComment struct{}
//
//		hidden int // unexported fields (and their comments) will not be shown in the documentation
//	}
//
// will be rendered as:
type Struct struct {
	// ExportedField can have a comment on the line above.
	ExportedField       int
	OtherExportedField  Links // or on the same line
	FieldWithoutComment struct{}

	hidden int // unexported fields (and their comments) will not be shown in the documentation
}

// You can mention bugs in the documentation. The syntax for that is like so:
//
//	// BUG(username): Some information.
//	// Some more information.
//
// This creates a section for bugs where each bug block is shown.

// You can use words other than "BUG", like "TODO". General syntax:
// MARKER(uid), with MARKER at least 2 chars (all caps) and uid at least 1 char.
//
// By default, only BUG notes are shown. If you run a godoc server locally, you
// can control that with the -notes argument. For example, -notes="BUG|TODO".
type Bugs doc.Note

// BUG(amit): This is an example bug.
// See the [Bugs] section.

// TODO(oliverpool): This is an example todo (hidden from godoc, unless run with -notes="BUG|TODO").
// See the [Bugs] section.

// Indented lines that start with dashes, stars, or pluses (or
// [other unicode characters]) create lists.
//
// For example:
//
//	This is my list
//	 - item 1
//	 - item 2
//
// Results in:
//
// This is my list
//   - item 1
//   - item 2
//
// [other unicode characters]: https://go.dev/doc/comment#lists
type Lists int

// Interface are rendered like Struct:
//   - exported methods are shown (with comments)
//   - unexported methods are hidden
type Interface interface {
	Method()
}

// You can mark a function or type as deprecated by adding a line that starts
// with "Deprecated:".
// See [Deprecated].
type Deprecation int

// An example deprecated type.
//
// Deprecated: use the new thing instead. See [Deprecation].
type Deprecated int
