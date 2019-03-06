package main

import "fmt"

func local(){
	a := 10
	fmt.Println("a=",a)
}
func main() {
	//定义在大括号里的就是局部变量
	//执行到定义变量那句话，才开始分配空间，离开作用域自动释放
	//作用域,变量其作用范围

	//a = 111

//	{
//		i := 10
//		fmt.Println("i=",i)
//	}
//	i = 111
//	fmt.Println("i=",i)
//	if flag :=1;flag==1{
//		fmt.Println("flag=",flag)
//	}
//	flag = 4
}
