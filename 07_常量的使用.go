package main

import "fmt"

func main(){
	const a int = 10
	//a = 20 常量不允许修改
	fmt.Println("a = ",a)

	const b = 10.1  //没有使用：=
	fmt.Printf("b type is %T\n",b)
	fmt.Println("b = ",b)

}
