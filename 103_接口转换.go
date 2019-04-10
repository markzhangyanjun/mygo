package main

import "fmt"

type Humaner3 interface {//子集
	sayhai()
}


type Personer3 interface{ //超集
	Humaner3
	sing(lrc string)
}


type Students struct {
	name string
	age  int
}

//Students实现了sayhi方法
func (tmp *Students) sayhai(){
	fmt.Printf("Students[%s,%d] sayhai\n",tmp.name,tmp.age)
}

func (tmp *Students)sing(lrc string){
	fmt.Printf("Students[%s,%d] sing\n",tmp.name,tmp.age)
}

func main() {
	//只要实现了此接口方法的类型，那么这个类型的变量（接受者类型）就可以给i赋值

	//定义一个接口变量
	//var i Humaner3
	//s := Students{"zyj",18}
	//i = &s
	//i.sayhai()

	//超集可以转换为子集，反过来不可以(子集不可以转换为超集)
	var ipro Personer3 //超集
	var ihum Humaner3 //子集

	//ipro = ihum //Humaner3 does not implement Personer3 (missing sing method)

	ihum = ipro
	s1 :=Students{"maek",18}
	ihum = &s1
	ihum.sayhai()


}