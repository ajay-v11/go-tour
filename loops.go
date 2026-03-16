package main

import "math"
import "fmt"

// Go weirdly has only for loops

func printSumInRange() int {

	sum := 1

	// This is how you write a normal for loop
	// for i := 1; i < 1000; i++ {
	// 	sum += i
	// }

	// This for statement has 3 parts, 1. init (i:=x) 2.conditional (i<y)  3.post statement(i++)
	// In go apparently the init and post statements are optional, so it kinda makes the "for" a while statement as well

	// This is how you write a while loop in go, which is actually a for loop
	for sum < 1000 {
		sum += sum
	}

	// Forever loop or infinte loop
	// for {
	// }
	return sum

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// In Go, fmt.Sprint converts values to a string.
