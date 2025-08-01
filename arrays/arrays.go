package main

import "fmt"

func main(){
fmt.Println("Arrays in Go")
 // dexlaring an array
 var favNums = [2]int{07, 27}
 fmt.Println(favNums)

 // declaring an array with :=

 favFood := [2]string{"Pizza", "Burger"}
 fmt.Println(favFood)

 // declaring an array without specifying the size

favPlace := [...]string{"My fav place is where , where my fav person is "}
fmt.Println(len(favPlace))

// declaring an array with a specific size

favNum := [5]int{1, 2, 3, 4, 5}
fmt.Println(favNum)

// accessing an array
fmt.Println(favNum[2])

// changing an array
favNum[2] = 10
fmt.Println(favNum)


}