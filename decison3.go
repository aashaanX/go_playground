package main

import "fmt"

func main(){
	var x interface{}
	x = 3.4
	fmt.Println("Switch  case is gonna execute")

	switch x.(type){
		case int: fmt.Println("is of type int")
		case string : fmt.Println("is of type string")
		case float64 : fmt.Println("is of type float")
		default: fmt.Println("is of type unknown")
	}
}
