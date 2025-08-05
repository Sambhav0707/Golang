package main


import "fmt"

// callback function is a function that is passed as an argument to another function
func truthOrDare(statement string ) string{
	if statement == "truth"{
		fmt.Println("I love you")
		return "I love you"
	}else{
		fmt.Println("I will do anything for you")
		return "I will do anything for you"
	}

}

func truthOrDare2(statement string , answer func(string) string){
	answer(statement)
}
func main() {

	statement := "Dare"
	truthOrDare2(statement , truthOrDare)

}