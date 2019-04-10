package main

import "fmt"

//定义接口类型
type Humaner interface{
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi()
}

type Student_u struct {
	name string
	id   int
}
//Student实现了此方法
func (tmp *Student_u) sayhi(){
	fmt.Printf("Student[%s,%d] sayhai\n",tmp.name,tmp.id)

}

type Teacher struct {
	addr string
	group string
}
//Teacher实现了此方法
func (tmp *Teacher)sayhi(){
	fmt.Printf("Teacher[%s,%s] sayhi\n",tmp.addr,tmp.group)

}

type Mystr string
//Mystr实现了此方法
func(tmp *Mystr)sayhi(){

	fmt.Printf("Mystr[%s] sayhi\n",*tmp)

}

func main(){

	//定义接口类型的变量
	var i Humaner
	//只要实现了此接口方法的类型，那么这个类型的变量（接受者类型）就可以给i赋值

	s := &Student_u{"mike",666}
	i = s
	i.sayhi()

	t := &Teacher{"mike","fsf"}
	i = t
	i.sayhi()


	var str Mystr="hello"
	i = &str
	i.sayhi()



}
