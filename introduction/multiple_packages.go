package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var num = math.Floor(2.75)
	var title = strings.ToTitle("My Go \nlearning course")
	fmt.Println(num)
	fmt.Println(title)
	// or we can simply use:
	fmt.Println(math.Floor(5.99))
}
