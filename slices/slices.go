package main

import "fmt"

func main() {
	fmt.Println("Slices")

	// declaring a slice
	slice := []int{}

	//appending to a slice

	slice = append(slice, 07, 27)
	fmt.Println(slice)

	// declaring a slice with a specific size
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice2), cap(slice2))

	// declaring a slice with a specific size and capacity
	slice3 := make([]int, 5, 7)
	fmt.Println(len(slice3), cap(slice3))

	// creating a slice from an array
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice4 := arr[1:4]
	fmt.Println(slice4)

	slice5 := arr[3:4]
	fmt.Println(slice5)


	// printing the length and capacity of the slice along with the slice
	slice6 := [10]int{}
	for i := 0; i < len(slice6); i++ {
		slice6[i] = i + 1
	}
	fmt.Println(slice6)

	
}
