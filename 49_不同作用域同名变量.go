package main

import "fmt"

var a byte//全局变量
//1.不同作用域允许定义同名变量
//2.使用变量的原则，就近原则
func main() {
	var a int //局部变量
	fmt.Printf("1 %T\n",a)
	{
		var a float64
		fmt.Printf("2 %T\n",a)
	}
	tes()
}

func tes(){
	fmt.Printf("3 %T",a)
}