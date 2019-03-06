package main

import "fmt"

//定义在函数外面的变量是全局变量
//全局变量在任何地方都可以使用
func all(){
	fmt.Println("a = ",a)
}
var a int
func main() {
	a = 10
	fmt.Println("a = ",a)
	all()
}