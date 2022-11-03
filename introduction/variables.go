package main

import "fmt"

func main() {
	var quantity = 10 // because we assigned value, we don't need int type
	fmt.Println(quantity)
	var num int // here we need to put int type, def value is 0
	num = 100
	fmt.Println(num)
	var name = "Anna"
	var phrase = "Hello " + name
	fmt.Println(phrase)
	length, width := 1, 12 // if vars are the same type, you can assign them like these
	fmt.Println(length, width)

	// you're not able to change type of variable
	var flatNum = 100
	// var flatNum = "blabla" it throws an err, because int variable isn't able to convert to string
	fmt.Println(flatNum)

	// empty def values
	var intNum int
	var floatNum float64
	var stringName string
	var boolVar bool
	// stringName return empty space by def
	fmt.Println(intNum, floatNum, stringName, boolVar)
}
