package main

import "fmt"

func main(){
	var x int=10
	var y int 
	
	numbers := [6]int{1,2,3,4,5}
	
	for y:=0; y<10;y++ {
		fmt.Printf("value of y:%d\n",y)
	}

	for y<x {
		y++
		fmt.Printf("value of y:%d\n",y)
	}

	for i,a := range numbers{
		fmt.Printf("value of x = %d at %d\n",a,i)
	}
}	