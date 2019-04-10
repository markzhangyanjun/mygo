package main

import "fmt"

type Person_w struct{
	name string
	sex  byte
	age  int
}

//有个学生，继承Person字段，成员和方法都继承了
type Student_a struct{
	Person_w
	id int
	addr string
}


func (tmp *Person_w)PrintInfo_a(){

	fmt.Printf("name=%s,sex=%c,age=%d\n",tmp.name,tmp.sex,tmp.age)

}


func main(){

	s := Student_a{Person_w{"mike",'f',17},666,"bj"}
	s.PrintInfo_a()


}

