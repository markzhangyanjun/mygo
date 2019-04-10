package main

import "fmt"

type Humaner2 interface{
	sayhi2()
}

type Personer interface {
	Humaner2 //匿名字段，继承了方法sayhi()
	sing(lrc string)
}

type Mark struct{
	name string
	age int
}
//Mark实现了此方法
func(tmp *Mark)sayhi2(){

	fmt.Printf("Student[%s,%d] sayhi2\n",tmp.name,tmp.age)

}
//Mark实现了此方法
func (tmp *Mark)sing(lrc string){
	fmt.Println("Mark在唱歌：",lrc)
}


func main(){

	//定义一个接口类型的变量
	var i Personer
	s := &Mark{"name",666}
    //只要实现了此接口方法的类型，那么这个类型的变量（接受者类型）就可以给i赋值
	i = s
	i.sayhi2()//继承过来的方法
	i.sing("hohoho")




}