package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main(){
	intMap:=make(map[string]int)

	intMap["one"]=1
	intMap["two"]=2
	intMap["three"]=3
	intMap["four"]=4
	intMap["five"]=5

	for key,value:= range intMap{
		// maps are not sorted even we put sorted
		// info in them
		fmt.Println(key,value)
	}
	

	// delete a key,e.g "four" 
	// delete(intMap,"four")

	el,ok:=intMap["four"]
	if ok{
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}

	// change the value of a given key
	intMap["two"]=4

}