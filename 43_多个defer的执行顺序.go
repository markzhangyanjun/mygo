package main

import "fmt"

func my(x int)  {
	res := 100/x
	fmt.Println("res=",res)
}
//先进后出
func main() {
	defer fmt.Println("aaaaa")
	defer fmt.Println("bbbbb")
	//调用一个函数。导致内存出问题
	defer my(0)
	defer fmt.Println("ccccc")
}
