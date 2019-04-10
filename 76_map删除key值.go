package main

import "fmt"

func main(){
	m := map[int]string{1:"mike",2:"zyj",3:"www"}
	fmt.Println("m=",m)

	//删除key为1	的内容
	delete(m,1)
	fmt.Println("m=",m)



}
