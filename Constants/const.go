package main

import "fmt"

func main() {
	// to create a constants in Go, we use the const keyword
	const pi = 3.14     // this is the way to create a constant in Go without specifying the type
	const Love int = 27 // this is the way to create a constant in Go with specifying the type
	fmt.Println("Love is with ", Love)
	fmt.Println("Pi is ", pi)
	// we can't change the value of a constant in Go after initialising

	//there is way to creat multiple constants at once

	const (
		love = 27
		Me   = 7
	)
	fmt.Println("Love is with ", love)
	fmt.Println("Me is with ", Me)

}
