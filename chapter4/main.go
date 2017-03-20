// A package comment is usually placed right above the package declaration
package main // package declaration is required to be the first line of code
import (
	"fmt"
	"math"
)

// Programs consist of keywords (25 predeclared), variables, constants,
// operators, types and functions. There are 36 predeclared identifiers.

//"factoring the keyword". Valid for import,const, var and type
// aliasing the package
//

// Suppose we have a thing (variable or function) called Thing (starts
// with T so it is exported) in a package pack1, then when pack1 is
// imported in the current package, Thing can be called with the usual
// dot-notation from OO-languages: pack1.Thing (The pack1. is necessary!)

// After the import statement, 0 or more const, var, types can be declared.
// These have package scope (global). Must be followed be one or more functions

// main must have parens but no arguments or return types. func main is the
// entry point of the program and exits immediately and successfully
// when main.main returns
// Code in functions is enclosed in braces

func main() { // The first brace must be on the same line as the declaration
	fmt.Println("Hello, world!!") // Calls the function Println from fmt

	// The var keyword automatically initializes the variable to the
	// zero value of its type
	var num int
	fmt.Println(num) // Prints 0

	fmt.Println(math.Pi)

	// Value types include: int, float, bool, string
	// They point directly to their value contained in memory

	// Reference types include: pointers, slices, maps and channels
	// Contains the address of the memory location where the value is stored

	var num2 int64
	num1 := int64(num)

	result := num1 + num2

	fmt.Println(result)

	// Execution of a function ends when a } or a return statementis reached
}
