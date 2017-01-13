package main

import (
	"fmt"
	"math"
)

func testFunction(name string){
	fmt.Printf("The name is : %s\n",name)
}

func testFunction2(num int) int{
	return num*2
}

func max(num1,num2 int) int{
	if(num1>num2){
		return num1
	}else{
		return num2
	}
}

func swap(x,y string) (string,string){
	return y,x
}

/* function using call by reference */
func swap2(p *int,q *int){
	var temp int = *p
	*p = *q
	*q = temp
}

/* Anonymous function */
func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}


func main(){
	
	/*Function on fly*/
	getSquareRoot := func(x float64) float64{
				return math.Sqrt(x)
			}

	var x,y int  = 3,10
	testFunction("Arjun")
	fmt.Println(testFunction2(2))
	fmt.Printf("The max value is :%d\n",max(9,9))
	var a,b = swap("jack","jill")
	fmt.Println("a:"+a+",b:"+b)
	swap2(&x,&y)
	fmt.Println(x,y)
	fmt.Println(getSquareRoot(10))

	/* implemetation of anonymous function */
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}