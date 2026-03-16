package main

import "fmt"

var c, python, java bool

var rust, zig, haskel = true, true, "idk"

func printVar() {
	var i int

	var x, y int = 1, 23

	// short variable declaration, walrus operator i guess
	// this can only be done inside a function & not outside, outside everything should be var
	kahanaKhazana := "I dont know why this came in my mind"

	fmt.Println(x, y, rust, zig, haskel)
	fmt.Print(kahanaKhazana)

	fmt.Println(i, c, python, java)
}
