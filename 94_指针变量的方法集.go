package main



import "fmt"

type Person_g struct{
	name string
	sex  byte
	age  int
}


//修改成员变量值


func (p Person_g)SetInfoValue(){

	fmt.Println("SetInfoValue")

}


func (p Person_g)SetInfoValuePoint(){

	fmt.Println("SetInfoValuePoint")


}


func main(){

	//加入结构体变量是一个指针变量，它能够调用方法，这些方法就是一个集合，简称方法集
	p :=&Person_g{"mike",'m',18}

	p.SetInfoValuePoint()
	(*p).SetInfoValuePoint() //把*p转换为p后调用
    //内部做的转换，先把指针p专成*p后再调用
	p.SetInfoValue()
	(*p).SetInfoValue()
}
