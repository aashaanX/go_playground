package main

import "fmt"

func f(num,a,b int) {
	if(num==0){
		return
	}
	fmt.Println(a)
	f(num-1,b,a+b)
}

func main(){
	f(8,0,1)
}
	