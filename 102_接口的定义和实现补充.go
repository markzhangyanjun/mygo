package main



import "fmt"

//定义接口类型
type Humaner1 interface{
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi1()
}

type Student_u1 struct {
	name string
	id   int
}
//Student实现了此方法
func (tmp *Student_u1) sayhi1(){
	fmt.Printf("Student[%s,%d] sayhai\n",tmp.name,tmp.id)

}

type Teacher1 struct {
	addr string
	group string
}
//Teacher实现了此方法
func (tmp *Teacher1)sayhi1(){
	fmt.Printf("Teacher[%s,%s] sayhi\n",tmp.addr,tmp.group)

}

type Mystr1 string
//Mystr实现了此方法
func(tmp *Mystr1)sayhi1(){

	fmt.Printf("Mystr[%s] sayhi\n",*tmp)

}

//定义一个普通函数，函数参数为接口类型
//只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner1){

	i.sayhi1()

}

func main(){

	s := &Student_u1{"mike",666}

	t := &Teacher1{"mike","fsf"}

	var str Mystr1="hello"

	//调用同一个函数，不同的表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner1,3)

	x[0]=s
	x[1] =t
	x[2] = &str
    //第一个返回下标，第二个返回下标的值
	for _ , i := range x{
		i.sayhi1()
	}




}




