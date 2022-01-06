package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

type Animal struct {
	Name string
	Sound string
	NumberOfLegs int
}

func (a Animal) Says(){
	fmt.Printf("A %s says %s", a.Name,a.Sound)
	fmt.Println("")
}

func (a Animal) HowManyLegs(){
	fmt.Printf("The %s has %d legs", a.Name,a.NumberOfLegs)
	fmt.Println("")
}
func main() {
	var dog Animal

	dog.Name = "dog"
	dog.Sound ="woof"
	dog.NumberOfLegs = 4
	dog.Says()
	dog.HowManyLegs()
	

	cat:=Animal{
		Name:"cat",
		Sound:"meaw",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()
}