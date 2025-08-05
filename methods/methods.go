package main


import "fmt"


type Person struct {
	name string
	age int 
}
// value receiver method what it does is it takes a copy of the struct and then it modifies the copy and then it returns the copy
func (p Person) sayHello() {

	fmt.Println("Hello , my name is ", p.name , "and my age is ", p.age)
}

// pointer receiver method what it does is it takes a pointer to the struct and then it modifies the struct and then it returns the pointer

func (p *Person) happyBirthday() {
	p.age++
}




func main() {
	var p1 Person 
	p1.name = "P__n__i"
	p1.age = 21

	p1.sayHello()
	p1.happyBirthday()
	p1.sayHello()
}