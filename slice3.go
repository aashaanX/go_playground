package main

import "fmt"

func main(){
	var numbers [] int
	fmt.Println(numbers)
	numbers = append(numbers,0)
	fmt.Println(numbers,len(numbers),cap(numbers))
	numbers = append(numbers,1,2,3,4,5,6,7,8)
	fmt.Println(numbers,len(numbers),cap(numbers))
	numbers1 := make([] int,len(numbers),len(numbers)*2-1) 
	copy(numbers1,numbers)
	fmt.Println(numbers1,len(numbers1),cap(numbers1))
}