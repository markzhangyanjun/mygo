package main

import "fmt"

func myfunc(temp ...int){
	for _ ,data := range temp{
		fmt.Println("data= ",data)
	}
}

func myfunc01(temp ...int){
	for _ ,data := range temp{
		fmt.Println("data= ",data)
	}
}
func myfunc02(args ...int){
	//全部元素传递给myfunc
	myfunc(args...)
	myfunc01(args[2:]...)
}


func main(){
	myfunc02(1,2,3,4)
}
