package main

import "fmt"

func main() {
	var a int
	fmt.Println("a=",a)

	var b ,c int
	fmt.Println("b=",b,"c=",c)

	a = 10
	fmt.Println("a=",a)
	fmt.Printf("a type is %T\n",a)

	//变量的初始化
	var e int = 20  //初始化
	var f = 20   //赋值 先声明后赋值
	fmt.Println("e=",e,"f= ",f)


	//自动推导类型，必需初始化，通过初始化的值确类型
	h := 30
	//%T打印变量所属的类型
	fmt.Printf("c type is %T\n",h)
}
