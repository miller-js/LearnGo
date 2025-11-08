package main

import "fmt"

func runVariables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Print(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Print(e)

	f := "apple"
	fmt.Println(f)
}
