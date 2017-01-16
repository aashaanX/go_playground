package main

import "fmt"

func main(){
	var i int =10
	var p *int
	var p2 *int
	p = &i
	fmt.Printf("Value of i : %d\n",i)
	fmt.Printf("Value of p : %x\n",p)
	fmt.Printf("Value of *p : %d\n",*p)
	fmt.Printf("Value of p2 : %x\n",p2)
	/*fmt.Printf("Value of *p2 : %d\n",*p2)*/
}