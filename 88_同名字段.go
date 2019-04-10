package main

import "fmt"

type Person_b struct {
	name string
	sex  byte
	age int
}


type Studentb struct {
	Person_b //只有类型，没有字段，匿名字段，继承了person的成员
	id   int
	addr string
	name string   //和Person_b同名
}


func main(){

	var s Studentb
    //默认规则，如果能在本地作用域找到此成员，就操作此成员，如果没有找到，找继承成员
	s.name = "mike" //操作的是Studentb的name 还是Personb的name
	s.sex = 'm'
	s.addr = "bj"
	s.age =18
	fmt.Printf("s=%+v\n",s)

	//显示调用
	s.Person_b.name = "yoyou"
	fmt.Printf("s=%+v\n",s)


}