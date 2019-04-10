package main

import "fmt"

type Personf struct{
	name string
	sex  byte
	age  int
}


//修改成员变量值

//接受者为普通变量,非指针，值语义，就是拷贝
func (p Personf)SetInfob(n string,s byte,a int){

	p.name = n
	p.sex = s
	p.age = a

	fmt.Printf("值语义：&p=%p\n",&p)

}

//接受者为指针变量，引用传递
func (p *Personf)SetInfom(n string,s byte,a int){

	p.name = n
	p.sex = s
	p.age = a

	fmt.Printf("引用语义：p=%p\n",p)


}


func(tmp Personf) PrintInfob(){
	fmt.Println("tmp=",tmp)
}
func main(){

	s1 :=Personf{"go",'h',18}
	fmt.Println("s1=",s1)
	fmt.Printf("&s1=%p\n",&s1)

	//值语义
	s1.SetInfob("mark",'f',18)
	fmt.Println("s1=",s1)

	//引用语义：接可以传递值也可以传递引用
	s1.SetInfom("mark",'f',18)
	fmt.Println("s1=",s1)

	(&s1).SetInfom("mark",'f',18)
	fmt.Println("s1=",s1)

}