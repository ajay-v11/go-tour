package main

import (
	"fmt"
	"math"
)

func typeConversion() {
	// In go type conversion should be done explicitly
	var x, y int = 3, 4

	var f float64 = math.Sqrt(float64(x*y + y*y))

	var z uint = uint(f)

	fmt.Println(x, y, z, f)
}

func typeInference() {
	// using this operator the types are infered based on the value
	v := 324.234 + 23i

	fmt.Printf("v is of type %T\n", v)
}
