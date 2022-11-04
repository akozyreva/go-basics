package main

import (
	"fmt"
	"reflect"
)

func main() {
	var price = 100
	fmt.Println("Price is", price, "dollars.")
	var taxRate = 0.08
	fmt.Println(reflect.TypeOf(taxRate)) // and def type is float64!
	var tax = float64(price) * taxRate
	var total = price + int(tax)
	fmt.Println("Total cost is", total, "dollars.")
	var availableFunds = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budbet?", int(total) <= availableFunds)

}
