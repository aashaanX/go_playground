package main

import "fmt"

func main(){
	numbers:= [] int {1,2,3,4,5}
	fmt.Println(numbers)
	fmt.Println(numbers[0:3])
	fmt.Println(numbers[2:])
	fmt.Println(numbers[2:3])
	rev(numbers)
	fmt.Println(fact(5))
}

func fact(num int) int {
	if(num==1){
		return 1
	}
	return num*fact(num-1)
}

func rev(x [] int) [] int{
	if(len(x)==0){
		return nil
	}	
	rev(x[1:])
	fmt.Printf("%d\t",x[0])
	return nil
}