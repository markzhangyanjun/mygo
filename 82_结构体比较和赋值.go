package main

import "fmt"

//定义一个结构体
type Student4 struct{
	id int
	name string
	sex byte //字符类型
	age int
	addr string
}


func main(){

	s1 := Student4{1,"milke",'m',18,"bj"}
	s2 := Student4{1,"milke",'m',18,"bj"}
	s3 := Student4 {2,"milke",'m',18,"bj"}

	fmt.Println("s1==s2",s1==s2)
	fmt.Println("s1==s2",s1==s3)

	//同类型的2个结构体变量可以相互赋值
	var tmp Student4
	tmp = s3
	fmt.Println("tmp=",tmp)

}




