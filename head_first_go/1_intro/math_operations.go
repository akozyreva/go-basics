package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 1.2
	var width int = 2
	// fmt.Println("Area is", length*width) it doesn't work - different types
	// fmt.Println("length > width?", length > width) it doesn't work - different types
	fmt.Println("Area is", length*float64(width)) // and how it works, we change type on the fly
	fmt.Println(reflect.TypeOf(width))
}
