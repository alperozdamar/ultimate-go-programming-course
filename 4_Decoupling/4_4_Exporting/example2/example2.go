// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

func main() {

	// Create a variable of the unexported type and initialize the value to 10.
	//counter := counters.alertCounter(10) // Compiler ERROR !

	// ./example2_test.go:17: cannot refer to unexported name counters.alertCounter
	// ./example2_test.go:17: undefined: counters.alertCounter

	//	fmt.Printf("Counter: %d\n", counter)
}
