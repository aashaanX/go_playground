package main

import "fmt"

func main(){
	var num int = 0
	var ans int = 1
	if(num==0){
		ans = 1
	}else{
		for x:=1;x<=num;x++{
			ans = ans* x
		}
	}
	fmt.Printf("ans:%d\n",ans)
}