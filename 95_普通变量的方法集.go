package main



import "fmt"

type Person_m struct{
	name string
	sex  byte
	age  int
}


//修改成员变量值


func (p Person_m)SetInfoValue(){

	fmt.Println("SetInfoValue")

}


func (p Person_m)SetInfoValuePoint(){

	fmt.Println("SetInfoValuePoint")


}


func main(){


	p :=Person_m{"mike",'m',18}
    //内部会先把p转化为 &p 	再调用
	p.SetInfoValuePoint()

	p.SetInfoValue()
}

