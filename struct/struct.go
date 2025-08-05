package main

import "fmt"

type FavPerson struct {
	fav    Fav
	number int
	name   string
}

type Fav struct {
	food     string
	beverage string
	dessert  string
	favActor string
}

func main() {

	var favPerson FavPerson

	favPerson.name = "P__n__i"
	favPerson.number = 27
	favPerson.fav.beverage = "chai"
	favPerson.fav.dessert = "Kaju patisa"
	favPerson.fav.favActor = "R Madhavan"

	fmt.Println(favPerson)

}
