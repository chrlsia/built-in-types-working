package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main(){
	z:=addTwoNumber(2,4)
	fmt.Println(z)

}

func addTwoNumber(x, y int) (sum int) {
	sum = x+y
	return
}