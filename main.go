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

import (
	"fmt"
	"testing"
)


/////////////////// Way More Examples than We Need //////////////////

/////////////////// The Original Type //////////////////

// Since this was a question specifically about implementation of
// Generics to maps in Go, here are some examples. I realize this
// much more than required, but I think examples help to clear
// up misconceptions.

// types based on the original interface-based type
type (
	// original interface{} map from the
	// Stack Overflow question. This is the
	// 'value' for the map example.
	//
	// used with:
	//  mapping := map[string]IExample[any]{
	//       // .......
	//  }
	IExample[T any] interface{ ExampleFunc(ex T) T }

	// A type for the entire map, instead of just
	// the 'value' used in the original example.
	// But ... is this the way to go? Is this clear
	// and does it communicate the functionality?
	IExampleMap[T any] map[string]interface{ ExampleFunc(ex T) T }

	// examples of IExampleMap instantiated for various types
	IExampleMapAny   = IExampleMap[any]
	IExampleMapInt   = IExampleMap[int]
	IExampleMapTests = IExampleMap[struct { // type for table based testing
		name    string
		fn      interface{}
		in      []interface{}
		want    []interface{}
		wantErr bool
	}]

	// It is getting somewhat cluttered ... perhaps it is more clear to define the interface first ...
	Exampler[T any] interface {
		ExampleFunc(ex T) T
	}

	Funcer[T any] func(ex T) T

	// AnyMap is an extremely generic definition, the
	// map type can be defined in any possible combination
	// of key and value types, so long as the 'key' type
	// is comparable.
	AnyMap[K comparable, V any] map[K]V

	// IntMap is a generic map with int keys
	IntMap[V any] AnyMap[string, V]

	// StringMap is a generic map with string keys.
	StringMap[V any] AnyMap[string, V]

	/////////////////// Best type definition for interface map //////////////////

	// This is probably the most clear and concise definition
	// for the original map type: string keys, Exampler values,
	// any type may be instantiated in the Exampler
	//
	ExamplerMap[T any] StringMap[Funcer[T]]

	// a few other ExamplerMaps instantiated for various types as examples
	IExampleFloat = ExamplerMap[float64]
	IExampleBool  = ExamplerMap[bool]
)

/////////////////// A Better Choice? //////////////////

// Based on the implicit goals in the original
// example, a function may be a better 'value'
// than an interface.
type (
	// type based on function signature instead
	// of interface implementation
	ExampleFunc[T any] func(ex T) T

	// using a function signature as the value
	FuncMap[T any] StringMap[ExampleFunc[T]]

	/////////////////// Best type definition for function map //////////////////

	// final type instantiated with 'any'
	ExampleFuncMap FuncMap[any]

	// EE[T any] = FuncMap[T]

	// example of a function map with multiple types of keys
	// and function argument allowed by the constraints
	ExampleFuncKeyMap[K comparable, V any] AnyMap[K, FuncMap[V]]
)

// Perhaps ... a useful set of defintions for running generic
// table-based tests.
type (
	TestRunner[O Ordered] func(t *testing.T, in ...any) (got O)

	Test[O Ordered] struct {
		// name    string // use map key - this also enforces the requirement to have unique names for tests ;)
		got     TestRunner[O]
		in      []any
		want    O
		wantErr bool
	}

	// TestMap is a map of table-based test objects
	TestMap[K comparable, V Ordered] AnyMap[K, Test[V]]
)

// CreateSampleInterfaceMap returns an example
// interface map with several sample values.
func CreateSampleInterfaceMap() IExampleMapAny {
	return IExampleMapAny{
		"any1": anyThing{"stuff"},
		"any2": anyThing{42},
		"any3": anyThing{"more different stuff"},
		"any4": anyThing{true},
	}
}

// CreateSampleFunctionMap returns an example map with
// several sample values.
func CreateSampleFunctionMap() ExampleFuncMap {
	// Mapping2 maps string names to examples. The examples
	// are stored as functions with the signature
	// 		func(ex T) T
	// which is instantiated to
	// 		func(ex any) any
	// in this example
	return ExampleFuncMap{

		// custom function
		"any": MyFuncAny,

		// original interface
		"a": anyThing{"stuff"}.ExampleFunc,

		// inline function
		"fake": func(ex any) any { return "fake" },
	}

	/*  TODO ...
	for k, v := range Mapping2 {
	// calling v with the zero value of the type for this
	// example; this data could be stored in a struct,
	// sent on a channel, stored in a file ...
	fmt.Printf("%v: %v\n", k, v(nil))
	*/
}

func main() {

	IExample := CreateSampleInterfaceMap()

	// add another example to Mapping
	IExample["any4"] = anyThing{"even more stuff"}

	FuncExample := CreateSampleFunctionMap()

	FuncExample["late addition"] = MyFuncAny(ex any) any { return "stuff" }

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
	// These are found in the main_test.go file
	// in the GitHub repo so they can be run
	// from go test in the standard fashion:
	PrintIExample(IExample)
	PrintFuncExample()

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

// PrintIExample actually can print any map, but is used
// here for the example type IExampleMap.
func PrintIExample[K comparable, V any](m map[K]V) {
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}
func PrintFuncExample() {
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

// some useful constraints - incomplete,  but good for example
// (e.g. leaves out complex numbers and uintptr, as well as
// ignoring structs, sequences, and sets that may be comparable
// based on their fields)
// as of this writing (4/01/22), the standard library 'constraints' package
// containing most of these constraints appears to be unavailable
type (
	Ordered interface{ Number | ~string }
	Number  interface{ IntType | UintType | FloatType }
	IntType interface {
		int | int8 | int16 | int32 | int64
	}
	UintType interface {
		uint | uint8 | uint16 | uint32 | uint64
	}
	FloatType interface{ float32 | float64 }
)
