// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var i int
	var s string
	var f float64
	var b bool

	// Display the value of those variables.
	fmt.Printf("var i int \t %T [%v]\n", i, i)
	fmt.Printf("var s string \t %T [%v]\n", s, s)
	fmt.Printf("var f float64 \t %T [%v]\n", f, f)
	fmt.Printf("var b bool \t %T [%v]\n", b, b)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	ii := 10
	si := "hello"
	fi := 3.14
	bi := true

	// Display the value of those variables.
	fmt.Printf("ii := 10 \t %T [%v]\n", ii, ii)
	fmt.Printf("si := \"hello\" \t %T [%v]\n", si, si)
	fmt.Printf("fi :=  3.14 \t %T [%v]\n", fi, fi)
	fmt.Printf("bi := true \t %T [%v]\n", bi, bi)

	// Perform a type conversion.

	// Display the value of that variable.
	iii := int32(ii)
	fmt.Printf("iii:= 10 \t %T [%v]\n", iii, iii)

}
