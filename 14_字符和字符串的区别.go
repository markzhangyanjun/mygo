package main

import "fmt"

func main(){
	var ch byte
	var str string
	//字符
	// 1.单引号
	//2.字符元素往往只有一个字符，转译字符除外
	ch = 'a'
	fmt.Println("ch = ",ch)
	//字符串
	//1.双引号
	//2.字符串有1个或多个字符组成
	//3.字符串都是隐藏了一个结束符'\0'
	str = "a" //由 a 和'\0'组成一个字符串
	fmt.Println("str=",str)

	str = "hello word"
	fmt.Printf("str[0] = %c,str[1] = %c\n",str[0],str[1])
}
