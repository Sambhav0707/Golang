package main

import "fmt"



// defer is used to delay the execution of a function until the surrounding function returns

func printHello() {
	fmt.Println("Hello")
}
func main() {
	defer printHello()
	fmt.Println("World")
}


// defer is used to delay the execution of a function until the surrounding function returns

