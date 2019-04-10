package main

import "fmt"

type Person_a struct {
	name string
	sex  byte
	age int
}


type Student10 struct {
	Person_a //只有类型，没有字段，匿名字段，继承了person的成员
	id   int
	addr string
}


func main(){

	s1 := Student10{Person_a{name:"mike",sex:'m',age:18},1,"bj"}

	fmt.Printf("s1=%+v\n",s1)

	s1.name = "youyou"
	s1.sex = 'w'
	s1.age = 22
	s1.id = 666
	s1.addr = "sz"
	fmt.Printf("s1=%+v\n",s1)
	fmt.Println(s1.name,s1.sex,s1.age,s1.id,s1.addr)

	s1.Person_a = Person_a{"go",'m',10}
	fmt.Printf("s1=%+v\n",s1)
	fmt.Println(s1.name,s1.sex,s1.age,s1.id,s1.addr)


}
