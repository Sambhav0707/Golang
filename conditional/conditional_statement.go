package main


import "fmt"


func main(){


	favNum := 27


	for i:= 0 ; i < 100 ; i++{
		if(i == favNum){
			fmt.Println("Found my fav number")
			break
		}else{
			fmt.Println("Not my fav number")
		}
	}

	//switch statement
	number := 7;

	switch number{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	default:
		fmt.Println("Not a number")
	}

	
}