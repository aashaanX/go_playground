package main

import "fmt"

func main(){
	var a,b,c = 3,4.2,"foo"
	fmt.Println(a)
	fmt.Printf("Type of a is %T\n",a)
	fmt.Println(b)
	fmt.Printf("Type of b is %T\n",b)
	fmt.Println(c)
	fmt.Printf("Type of c is %T\n",c)
}