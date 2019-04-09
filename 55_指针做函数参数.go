package main

import "fmt"

func swap1(p1,p2 *int){

	fmt.Printf("*p1=%d,*p2=%d\n",*p1,*p2)

	*p1,*p2 = *p2,*p1
}
func main(){

	a, b := 10,20
	swap1(&a,&b)//传递地址

	fmt.Printf("a=%d,b=%d",a,b)
}
