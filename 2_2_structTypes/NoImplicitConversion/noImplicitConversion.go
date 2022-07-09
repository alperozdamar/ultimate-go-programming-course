// Sample program to show how to declare and initialize anonymous
// struct types.
package main

import "fmt"

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	var a alice
	var b bill

	//b = a // Can not use a (type Alice ) as the type Bill. Go does not allow implicit Type conversion.
	// Because historically it caused more problems than good.

	b = bill(a) // Explicitly Convert

	b = e2 // Compiler does not ask explicit conversion here!!!

	//a is a name type, e2 is a literal type. When type is named, there is going to be NO implicit conversion.
	//but if it is a literal type, we have flexibility on assignment. (implicit conversion)

	fmt.Println(b, a)

}
