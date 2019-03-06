package main

import "fmt"

// ...int 类型这样的类型 ， ...type 不定参数类型
//注意：不定参数，一定只能放在形参中的最后一个参数
func MyFunc1(args ...int){//传递的参数可以是0个或多个

	fmt.Println("len(args)=",len(args))

	for i := 0; i < len(args); i ++ {
		fmt.Printf("args[%d]=%d\n",i,args[i])
	}

	for _,data := range args{
		fmt.Printf("data=%d\n",data)
	}

}
//固定参数一定要传参，不定参数根据需要传参
func MyFunc2(a,b string,args ...int){

}

//错误演示
//func Myfunc3(args ...int,a sting){

//}

func main(){
	MyFunc1(1,2,3)

}
