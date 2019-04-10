package main

import "fmt"

//定义一个结构体
type Student3 struct{
	id int
	name string
	sex byte //字符类型
	age int
	addr string
}


func main(){
	//指针有合法指向后，猜操作成员
    //先定义一个普通的结构体变量
    var s Student3
    //再定义一个指针变量，保存s的地址
    var p1 *Student3
    p1 = &s
    //通过指针操作成员 p1.id 和（*p1).id 完全等价，只能使用 . 运算符
    p1.id = 18
	(*p1).name = "zyj"
	p1.age = 18
	p1.addr = "bj"

	fmt.Println("p1=",p1)


	//通过new申请一个结构体
	p2 :=new(Student3)
	p2.id = 18
	(*p2).name = "zyj"
	p2.age = 18
	p2.addr = "bj"

	fmt.Println("p2=",p2)

}

