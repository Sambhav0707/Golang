package main

import "fmt"

func addLoveNumbers(a int, b int) int {

	fmt.Println(a + b)
	return a + b
}

// function with return name

func addLoveNumbersWithReturn(a int, b int) (result int) {
	result = a + b
	fmt.Println(result)
	return
}

// function can return multiple values

func addLoveNumbersWithMultipleReturn(a int, b int) (result int, result2 int) {
	result = a + b
	result2 = a - b
	fmt.Println(result, result2)
	return
}

// recursive function
func recursiveFunction(a int) int {
	if a == 0 {
		return 7
	}
	fmt.Println(a)
	recursiveFunction(a - 1)
	return a

}

func main() {
	var a int = 7
	var b int = 27

	addLoveNumbers(a, b)
	addLoveNumbersWithReturn(a, b)
	_, result2 := addLoveNumbersWithMultipleReturn(27, 7)
	fmt.Println(result2)
	recursiveFunction(0)

}
