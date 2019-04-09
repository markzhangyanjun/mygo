package main

import (

	_ "calc"
	"fmt"
)



func init(){
	fmt.Println("this is main init")
}

func main(){
	fmt.Println("this is main")
	//a :=calc.Add(10,20)

	//fmt.Println("a=",a)
}
