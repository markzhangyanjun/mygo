package main

import "fmt"

func Add(a,b int)int{
	return a + b
}

func Minus(a,b int)int{
	return a - b
}

//函数也是一种数据类型
//FuncType 它是函数类型
type FuncType func(int,int)int //没有函数名字，没有{}
func main(){
	var res int
	res = Add(1,1)
	fmt.Println("res = ", res)
	//声明一个函数类型变量，变量名叫fTest
	var fTest FuncType
	fTest = Add //是便量就可以赋值
	res = fTest(10,20)
	fmt.Println("res = ", res)
}
