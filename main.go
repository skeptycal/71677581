// package main contains a Response to Stack Overflow question 71677581
//
// Stack Overflow Post: https://stackoverflow.com/q/71677581
//
// Go Playground version: https://go.dev/play/p/jR5MpX-Fopc
//
// This repo: https://github.com/skeptycal/71677581
//
// Copyright (c) 2022 Michael Treanor
//
// MIT License
//
// GitHub: https://github.com/skeptycal
//
// Twitter: https://twitter.com/skeptycal
package main

import "fmt"

// Copyright (c) 2022 Michael Treanor
// www.github.com/skeptycal/

type (
	// original interface{} map from the question:
	IExample[T any] interface {
		ExampleFunc(ex T) T
	}
	// single type
	IExampleMap[T any] map[string]interface{ ExampleFunc(ex T) T }

	// map type based on function signature instead of interface method
	Example[T any] func(ex T) T

	// single type
	ExampleMap[T any] map[string]func(ex T) T

	// 1Example instantiated for type any
	IExampleAny = IExample[any]

	// 1Example instantiated for type float64
	IExampleFloat = IExample[float64]

	// 1Example instantiated for type bool
	IExampleBool = IExample[bool]
)

// mapping maps string names to examples. The
// examples are instantiated with [any] and
// thus are filled in with interface objects
// that implement ExampleFunc(ex any) any
var Mapping = map[string]IExample[any]{
	"any1": anyThing{"stuff"},
	"any2": anyThing{"other stuff"},
	"any3": anyThing{"more different stuff"},
}

func ExampleMappingMain() {
	for k, v := range Mapping {
		fmt.Printf("%v: %v\n", k, v)
	}

	// Output:
	// any1: {stuff}
	// any2: {other stuff}
	// any3: {more different stuff}
}
func ExampleMapping2Main() {
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

func main() {

	// add another example to Mapping
	Mapping["any4"] = anyThing{"even more stuff"}

	// ************************************** error
	// uncomment the next line to see the error in action!
	// Mapping["float64"] = float64thing{42.0}

	// compile time error: InvalidIfaceAssign
	/*************  InvalidIfaceAssign  *************
	  cannot use (float64thing literal) (value of type float64thing) as IExample[any] value in map literal: float64thing does not implement IExample[any] (wrong type for method ExampleFunc)
	   		have ExampleFunc(ex float64) float64
	   		want ExampleFunc(ex any) any m
	  ************************************************/

	// examples:
	// var mappingFloat map[string]IExample[float64]
	// var mappingFloat map[string]IExampleFloat

	// var mapping3 map[string]IExample[any]

	// instantiated with int:
	// var intMap ExampleMap[int]

	/////////////////////////////////////// Tests
	// run from go test:
	ExampleMappingMain()
	ExampleMapping2Main()

}

// Mapping2 maps string names to examples. The examples
// are stored as functions with the signature
// 		func(ex T) T
// which is instantiated to
// 		func(ex any) any
// in this example
var Mapping2 = ExampleMap[any]{

	// custom function
	"any": MyFuncAny,

	// original interface
	"a": anyThing{"stuff"}.ExampleFunc,

	// inline function
	"fake": func(ex any) any { return "fake" },
}

// anyThing is a sample of some object
// that implements IExample
type anyThing struct{ any }

func (t anyThing) ExampleFunc(ex any) any { return t.any }

// intThing is a similar object that
// impements IExample ...
type intThing struct{ int }

func (t intThing) ExampleFunc(ex int) int { return t.int }

type float64thing struct{ float64 }

func (t float64thing) ExampleFunc(ex float64) float64 { return t.float64 }

func MyFuncAny(ex any) any { return "stuff" }
func MyFuncInt(ex int) int { return 42 }
