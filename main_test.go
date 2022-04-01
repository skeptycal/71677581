package main

import "fmt"

func ExampleMapping() {
	for k, v := range Mapping {
		fmt.Printf("%v: %v\n", k, v)
	}

	// Output:
	// any1: {stuff}
	// any2: {other stuff}
	// any3: {more different stuff}
}
func ExampleMapping2() {
	for k, v := range Mapping2 {
		// calling v with the zero value of the type for this
		// example; this data could be stored in a struct,
		// sent on a channel, stored in a file ...
		fmt.Printf("%v: %v\n", k, v(nil))
	}

	// Output:
	// any: stuff
	// a: stuff
	// fake: fake
}
