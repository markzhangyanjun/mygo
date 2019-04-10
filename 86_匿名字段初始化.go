package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age int
}


type Student9 struct {
	Person //只有类型，没有字段，匿名字段，继承了person的成员
	id   int
	addr string
}

func main(){
	//顺序初始化
	var s1 Student9 = Student9{Person{"mike",'m',18},1,"bj"}

	fmt.Println("s1=",s1)

	//自动推导类型
	s2 := Student9{Person{"mike",'m',18},1,"bj"}
	fmt.Println("s2=",s2)

	//%+v,显示更详细
	fmt.Printf("s2=%+v\n",s2)


	//指定成员初始化，没有初始化的成员自动赋值为0或为空
	s3 := Student9{id:1}
	fmt.Printf("s3=%+v\n",s3) //s3={Person:{name: sex:0 age:0} id:1 addr:}


	s4 := Student9{Person:Person{name:"mark"},id:1} //s4={Person:{name:mark sex:0 age:0} id:1 addr:}
	fmt.Printf("s4=%+v\n",s4)




}
