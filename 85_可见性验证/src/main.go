package main

import (
	"fmt"
	"test"
)



func main(){

	//包名.函数名
	test.MyFunc()

	//包名.结构体类型名

	var s test.Stu

	fmt.Println("s=",s)

	s.Id = 666
	fmt.Println("s=",s)

}
