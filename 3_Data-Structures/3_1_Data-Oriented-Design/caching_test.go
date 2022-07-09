// go test -run none -bench . -benchtime 3s
// Tests to show how Data Oriented Design matters.
/**
    The way the benchmark works is, the Go testing tool, find these benchmark functions,
and call them one at a time, and when it calls the benchmark function, it will set this b.N
variable to be one. All of the code that we want to benchmark must be inside this loop. Now
it's important that the code inside the loop is as accurate to the way it's being used in
production as possible, or the benchmark isn't necessarily going to be accurate. And, in
order to make benchmarks accurate, your machine also must be idle.
*/
package caching

import "testing"

var fa int

// Capture the time it takes to perform a link list traversal.
func BenchmarkLinkListTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}

	//Now, another thing you have to understand is when we run these benchmarks, the Go compiler is
	//going to recompile this code. The compiler has the ability to throw dead code away. In other
	//words, the compiler could identify that this function doesn't mutate anything, yet it returns
	//a value and we're not storing it. That means the compiler could say, I'm not gonna waste my
	//time calling this function, because it has no impact on the output of this program whatsoever.
	//So, notice that I'm performing the assignment to a local variable, and then to a global variable,
	//so the compiler has no chance to throw this code away. You don't want this blazing fast benchmark,
	//and you don't want it to be blazing fast because it didn't do anything.
	fa = a
}

// Capture the time it takes to perform a column traversal.
func BenchmarkColumnTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = ColumnTraverse()
	}

	fa = a
}

// Capture the time it takes to perform a row traversal.
func BenchmarkRowTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = RowTraverse()
	}

	fa = a
}
