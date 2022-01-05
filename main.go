package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main(){
	x:=10
	myFirstPointer := &x

	fmt.Println("x is",x)
	fmt.Println("My first pointer is",myFirstPointer)
	fmt.Println("Contents of x are",*myFirstPointer)

	*myFirstPointer= 15
	fmt.Println("x now is:",x)


}