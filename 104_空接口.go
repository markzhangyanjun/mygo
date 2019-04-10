package main

import "fmt"



func xxx(arg ...interface{}){
	//可以接受0到多个任意类型的参数

}

func main() {
	//空接口：万能类型可以保存任意类型的值

	var i interface{} = 1

	fmt.Println("i=",i)

	i = "abc"
	fmt.Println("i=",i)
}
