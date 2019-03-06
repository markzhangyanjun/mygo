package main

import "fmt"

func main(){
	var flag bool
	flag = true
	fmt.Printf("flag=%t\n",flag)
	//bool类型不能转换为 int
	// fmt.Printf("flag=%d\n",int(flag))
	//整型也不能转换为bool类型
	//flag = bool(1)
	var ch byte
	ch = 'a' //字符型本质上就是整型
	var t int
	t = int(ch)
	fmt.Println("t=",t)
}
