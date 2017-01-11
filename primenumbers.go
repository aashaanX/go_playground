package main

import "fmt"

func main(){
	for i:=2;i<100;i++{
		var flag int = 0
		for j:=2;j<=(i/2);j++{
			if((i%j)==0){
				flag = 1
			}
		}
		if(flag==0){
			fmt.Println(i)
		}
	}
}
			