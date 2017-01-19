package main

import "fmt"

func main(){
	var countryCapitalMap map[string] string
	countryCapitalMap = make(map[string] string)
	fmt.Println(countryCapitalMap)
	countryCapitalMap["france"] = "paris"
	countryCapitalMap["spain"] = "barcelona"
	countryCapitalMap["india"] = "new delhi"
	fmt.Println(countryCapitalMap)
	capital,ok := countryCapitalMap["france"]
	if ok{
		fmt.Println(capital)
	}else{
		fmt.Println("Doesn't Exist")
	}

}