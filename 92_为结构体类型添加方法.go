package main

import "fmt"

type Persone struct {
	name string
	sex  byte
	age int
}

//带有接受者的函数叫方法
func(tmp Persone)PrintInfo(){
	fmt.Println("tmp=",tmp)
}
//通过一个函数给成员赋值
func (p *Persone)SetInfo(n string,s byte,a int){

	p.name = n
	p.sex = s
	p.age = a

}




//	参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接⼝或指针。
type long_a int

func (tmp long_a)test_a(){

}
//type pointer *int
//pointer 为接受者类型，它本身不能为指针类型
//func (tmp pointer)test_b(){ //invalid receiver type pointer (pointer is a pointer type)

//}


//不支持重载方法，也就是说，不能定义名字相同但是不同参数的方法。

//只要接受者不一样，这个方法就算是同名也是不同的方法，不会出现重复定义函数的错误
type  char byte
func (tmp char)test_a(){

}
type long_b int
func (tmp long_b)test_a(){

}









func main(){

	//定义同时初始化
	p := Persone{"mark",'f',18}
	p.PrintInfo()

	//定义一个结构体变量
	p1 := Persone{}

	(&p1).SetInfo("zyj",'f',18)
	p1.PrintInfo()

	p2 := new(Persone)
	p2.SetInfo("zyj",'f',18)
	(*p2).PrintInfo()




}