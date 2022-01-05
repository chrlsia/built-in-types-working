package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main(){
	var animals []string
	animals=append(animals,"dog")
	animals=append(animals,"fish")
	animals=append(animals,"cat")
	animals=append(animals,"horse")

	fmt.Println(animals)

	for _,x:= range animals{
		fmt.Println(x)
	}

	fmt.Println("Element 0 is:",animals[0])
	fmt.Println("First 2 elements are:", animals[0:2])
	fmt.Println("All the slice",animals[:])
	fmt.Println("Elements from position 2 kai mexri telos", animals[2:])
	fmt.Println("Elements from start kai pare 3 elements", animals[:3])

	fmt.Println("The slice is", len(animals),"elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	animals = DeleteFromSlice(animals,1)
	sort.Strings(animals)// in place sort
	fmt.Println(animals)

}

func DeleteFromSlice(a []string, i int) []string {
	a[i]=a[len(a)-1]
	a[len(a)-1]=""
	a=a[:len(a)-1]
	return a
}