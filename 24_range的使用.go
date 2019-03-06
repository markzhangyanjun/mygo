package main

import "fmt"

func main(){
	str := "abcdef"
	//通过for打印每个字符
	for i := 0 ; i < len(str) ; i ++ {
		fmt.Printf("%T\n",str[i])
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}
	//迭代打印每个元素，默认返回2个值：一个是元素位置，一个是元素本身
	for j,data := range str {
		fmt.Printf("str[%d]=%c\n",j,data)
	}

	for j := range str { //第二个返回值，默认丢弃，返回元素的位置
		fmt.Printf("str[%d] ",j)
	}

	for j,_ := range str { //第二个返回值，默认丢弃，返回元素的位置
		fmt.Printf("str[%d] ",j)
	}

}
