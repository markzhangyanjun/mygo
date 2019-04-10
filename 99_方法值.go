package main

import "fmt"

type Person_o struct{
	name string
	sex  byte
	age  int
}


//修改成员变量值


func (p Person_o)SetInfoValue1(){

	fmt.Printf("SetInfoValue:%p,%v\n",&p,p)

}


func (p *Person_o)SetInfoValuePoint1(){

	fmt.Printf("SetInfoValuePoint:%p,%v\n",p,p)


}


func main(){


	p :=Person_o{"mike",'m',18}
	fmt.Printf("main:%p,%v\n",&p,p)
	p.SetInfoValuePoint1() //传统嗲用方式


	//保存方式入口
	pFunc :=p.SetInfoValuePoint1 //这个就是方法值，调用函数时，无需再传递接受者，隐藏了接受者
	pFunc() //等价于 p.SetInfoValuePoint1()

	vFunc := p.SetInfoValue1
	vFunc()






}


