package main

import "fmt"

func fibu(num,a,b int) {
	if(num==0){
		return
	}
	fmt.Println(a)
	fibu(num-1,b,a+b)
}

func main(){
	fibu(8,0,1)
}
	