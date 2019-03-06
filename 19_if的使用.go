package main

import "fmt"

func main(){
	s := "王思聪"

	if s == "王思聪"{//左括号和if在一行
		fmt.Println("哈哈哈")
	}
	//if 支持一个初始化语句，初始化语句和判断条件以分号分隔
	if a := 10 ; a==10{
		fmt.Println("a=10")
	}
}
