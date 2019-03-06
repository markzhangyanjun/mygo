package main

import "fmt"

func main(){
	var str1 string
	str1 = "abc"
	fmt.Println("str1=",str1)

	//自动推导类型
	str2 := "hello"
	fmt.Printf("str2 type is %T\n",str2)

	//内建函数len()可以测字符串长度
	fmt.Println("len(str2)=",len(str2))
}

