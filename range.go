package main

import "fmt"

func main(){
	numbers := [] int {1,2,3,4,5}
	fmt.Println(numbers)
	for i:= range(numbers){
		fmt.Println(numbers[i])
	}

	/*Create a MAP*/
	countryCapitalMap := map[string] string {"france":"paris","india":"delhi","srilanka":"colombo","usa":"newyork","south africa":"johanasberg"}

	for i:= range countryCapitalMap{
		fmt.Printf("capital of %s is %s\n",i,countryCapitalMap[i])
	}

	for country,capital := range countryCapitalMap{
		fmt.Println("The Capital of",country,"is",capital)
	}
}