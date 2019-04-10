package main

import "fmt"

//定义一个结构体
type Student7 struct{
	id int
	name string
	sex byte //字符类型
	age int
	addr string
}

func testh1(p *Student7){
	p.id = 666
	fmt.Println("test01:",*p)
}


func main(){
	s1 := Student7{1,"milke",'m',18,"bj"}

	testh1(&s1) //地址传递，形参可以改实参
	fmt.Println("main:",s1)
}
