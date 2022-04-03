package main

func ExampleMapping() {
	PrintIExample(CreateSampleInterfaceMap())

	// Output:
	// any1: {stuff}
	// any2: {other stuff}
	// any3: {more different stuff}
}
func ExampleMapping2() {
	PrintIExample(CreateSampleFunctionMap())

	// Output:
	// any: stuff
	// a: stuff
	// fake: fake
}
