package main

import "fmt"

type mystra string //自定义类形

type Persond struct {
	name string
	sex  byte
	age int
}


type Studentd struct {
	*Persond //指针类型
	int     //基础类型的匿名字段
	mystra

}


func main(){
	s := Studentd{&Persond{"mike",'m',18},666,"hello"}
	fmt.Println("s=",s) //s= {0xc00000a080 666 hello}
	fmt.Println(s.name,s.sex,s.age,s.int,s.mystra) //mike 109 18 666 hello

	//先定义变量
	var s2 Studentd

	s2.Persond = new(Persond) //分配空间

	s2.name = "yoyou"
	s2.sex = 'm'
	s2.age = 22
	s2.int = 3
	s2.mystra ="bj"
	fmt.Println("s2=",s2)





}
