// Declare three 2_1_variables that are initialized to their zero value and three
// declared with a literal value. Declare 2_1_variables of type string, int and
// bool. Display the values of those 2_1_variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

func main() {

	// Declare 2_1_variables that are set to their zero value.
	var age int
	var name string
	var legal bool

	// Display the value of those 2_1_variables.
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	// Declare 2_1_variables and initialize.
	// Using the short variable declaration operator.
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	// Display the value of those 2_1_variables.
	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)
}
