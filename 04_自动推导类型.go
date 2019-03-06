package main

import "fmt"
func main(){
	var a int
	 a = 10
	 fmt.Println("a=",a)
     //:= 自动推导类型，先声明b，再给b赋值
	 b := 10
	 fmt.Println("b=",b)

	 //b :=20 //前面已经有变量，不能再新建一个变量

	 b = 30
	 fmt.Println("b=",b)

}
