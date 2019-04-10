package main

import "fmt"

//定义一个结构体
type Student5 struct{
	id int
	name string
	sex byte //字符类型
	age int
	addr string
}

func testh(s Student5){
	s.id = 666
	fmt.Println("test01:",s)
}


func main(){
	s1 := Student5{1,"milke",'m',18,"bj"}

	testh(s1) //值传递，形参无法改实参
	fmt.Println("main:",s1)
}
