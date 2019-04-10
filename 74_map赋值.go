package main

import "fmt"

func main(){
	m1 := map[int]string{1:"helloi",2:"go"}
	fmt.Println("m1=",m1)
	//赋值，如果key已经存在key值，修改内容
	m1[1] = "world"
	fmt.Println("m1=",m1)

	m1[3]= "youy" //追加，map底层自动扩容，和append类似
	fmt.Println("m1=",m1)




}
