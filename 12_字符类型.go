package main

import "fmt"

func main(){
	var ch byte
	ch = 97
	fmt.Println("ch=",ch)
	//格式化输出，%c以字符方式输出，%d以整型方式输出
	fmt.Printf("%c,%d\n",ch,ch)

	ch = 'a' //字符是单引号的
	fmt.Printf("%c,%d",ch,ch)
}
