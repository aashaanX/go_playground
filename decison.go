package main

import "fmt"

func main(){
	var a int = 20

	if (a>20){
		fmt.Println("a is greater than 20")
	}else if(a<20){
		fmt.Println("a is less than 20")
	}else{
		fmt.Println("a is equal to 20")
	}
	fmt.Printf("a value : %d\n",a)
}
