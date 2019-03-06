package main

import "fmt"

func main(){
	//a := 10
	//b := 20
	//c := 30
	a ,b := 10,20
	//交换两个变量的值
	var tem int
	tem = a
	a = b
	b = tem
	fmt.Printf("a = %d,b = %d\n",a,b)

	i , j := 10 ,20
	i , j = j ,i
	fmt.Println("i=",i,"j=",j)

    //_匿名变量，丢弃数据不处理,_匿名变量配合函数返回值使用才有优势
    _ , m := i ,j
    fmt.Println("m=",m)

}
