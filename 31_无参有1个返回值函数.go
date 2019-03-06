package main

import "fmt"

func test()int{
	return 666
}

func test1()(rest int){
	return 777
}
//常用写法
func test2()(rest int){
	rest = 888
	return
}

func main(){
	var a int
	a = test()
	fmt.Println("a=",a)

	b := test()
	fmt.Println("b=",b)

	c := test1()
	fmt.Println("c=",c)

	d := test2()
	fmt.Println("d=",d)
}
