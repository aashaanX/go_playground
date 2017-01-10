package main

import "fmt"

func main(){
	var  mark int = 70
	var grade string

	switch mark{
		case 90: grade = "A"
		case 80: grade = "B"
		case 70: grade = "C"
	}
	switch grade{
		case "A": fmt.Println("A")
		case "B": fmt.Println("B")
		case "C": fmt.Println("C")
	}
}
