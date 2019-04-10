package main

import "fmt"

type Person_x struct{
	name string
	sex  byte
	age  int
}

//有个学生，继承Person字段，成员和方法都继承了
type Student_b struct{
	Person_x
	id int
	addr string
}

func (tmp *Person_x)PrintInfo_b(){

	fmt.Printf("name=%s,sex=%c,age=%d\n",tmp.name,tmp.sex,tmp.age)

}

//Student也实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (tmp *Student_b)PrintInfo_b(){

	fmt.Println("Student:tmp=",tmp)

}






func main(){

	s := Student_b{Person_x{"mike",'f',17},666,"bj"}
	//就近原则，先找本作用域的方法，找不到再找继承的方法

	s.PrintInfo_b() //到底是调用Person还是Student ，结论是Student

	//显示调用继承的方法
	s.Person_x.PrintInfo_b()


}


