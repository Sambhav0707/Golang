package main

import "fmt"

func main() {
	/*
		There are two ways to declare variables in go . The first way is to use the var keyword.
		the second way is to use the short variable declaration syntax. that is ":=" operator.
	*/

	var favNumber int = 27 // declare and initialize a variable using var keyword 
	fmt.Println("My favourite number is ", favNumber)

	var favMonth = "FEBUARY" // delare and initialize a variable using var keyword without specifying the type
	fmt.Println("My favourite month is ", favMonth)

	favSweet := "Kaju Pateesa" // declare and initialize a variable using short variable declaration syntax
	fmt.Println("My favourite sweet is ", favSweet)



}