package main

import "fmt"

func test0(a int){
	if a==1{//函数终止的条件，非常重要
		fmt.Println("a=",a)
		return //终止函数调用
	}
	//函数调用自身
	test0(a -1)
	fmt.Println("a=",a)
}


func main(){
	test0(3)
	fmt.Println("main")

}