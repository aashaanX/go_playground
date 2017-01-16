package main

import "fmt"

func main(){
	var i int = 10
	ptr := &i
	ptr2 := &ptr
	fmt.Println(i)
	fmt.Println(ptr)
	fmt.Println(ptr2)
	fmt.Println(*ptr)
	fmt.Println(*ptr2)
	fmt.Println(**ptr2)
}