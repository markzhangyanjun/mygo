package main

import "fmt"

//定义一个结构体
type Student2 struct{
	id int
	name string
	sex byte //字符类型
	age int
	addr string
}


func main(){

	//定义一个普通变量
	var s Student2

	//操作成员，需要使用点（.）运算符

	s.id = 1
	s.name = "maki"
	s.sex = 'm'
	s.age =18
	s.addr = "bj"
	fmt.Println("s=",s)


}
