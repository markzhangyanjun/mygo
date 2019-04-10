package main

import "fmt"

type mystr string //自定义类形

type Personc struct {
	name string
	sex  byte
	age int
}


type Studentc struct {
	Personc //结构体匿名字段，
	int     //基础类型的匿名字段
	mystr

}


func main(){
	s := Studentc{Personc{"mike",'m',18},666,"hello"}
	fmt.Println("s=",s)

	s.name = "go"
	s.age = 18
	s.sex = 'f'
	s.int = 777
	s.mystr = "hehhf"
	fmt.Println(s.name ,s.age,s.sex,s.sex,s.int,s.mystr )
	fmt.Println(s.Personc,s.int,s.mystr)

	s.Personc=Personc{"go",'m',888}
	fmt.Println(s.Personc,s.int,s.mystr)

}


