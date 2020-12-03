package main

import "fmt"

// declare a struct
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Tea",
		contactInfo: contactInfo{
			email:   "jimtea@gmail.com",
			zipCode: 9400,
		},
	}
	// initialize a struct

	// first way to initialize a struct
	// alex := person{"Alex", "Anderson"}

	// second way to initialize a struct
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }

	// third way to initialize a struct
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)

	// fmt.Printf("%+v", jim)

	// define a pointer to a person called jim
	// converts a value to a pointer
	jimPointer := &jim
	// fmt.Println(jimPointer)
	jimPointer.updateName("Jordan")
	jim.print()
}

// the receiver type of the function is a pointer to a person
// i.e. only a pointer to a person could call this function
func (pointerToPerson *person) updateName(name string) {
	// converts a pointer to an actual value like struct
	(*pointerToPerson).lastName = name
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
