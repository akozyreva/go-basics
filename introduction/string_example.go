package main

import "fmt"

func main() {
	fmt.Println("Hello Go"[0])         // it shows ASCII bytes code for char
	fmt.Println(string("Hello Go"[0])) // it shows first letter

	// for symbols there will be ''
	fmt.Println('a')  // and it shows bytes code, because Go saves symbols so
	fmt.Println('\n') // the same - shows byte of symbol
	fmt.Println("a")  // and it's row and shows a
}
