package main

import "fmt"

func main(){
	a := 10
	str := "mark"
	//匿名函数，没有函数名字,函数定义，没有调用
	f1 := func(){//：=自动推导
		fmt.Println("a=",a)
		fmt.Println("str=",str)
	}
	f1()

	type functype func()
	//声明变量
	var f2 functype
	f2 = f1
	f2()

	//定义匿名函数同时调用
	func(){
		fmt.Printf("a=%d,str=%s\n",a,str)
	}()//后面的（）代表调用此匿名函数

	//带参数的匿名函数
	f3 :=func(i,j int){
		fmt.Println("i=",i,"j=",j)
	}
	f3(1,2)
	//带参数匿名函数同时调用
	func(i,j int){
		fmt.Println("i=",i,"j=",j)
	}(10,20)

	//匿名函数有返回值
	max , min  := func(i,j int)(max ,min int){
		if i > j{
			max = i
			min = j
		}else{
			max = j
			min = i
		}
		return
	}(111,222)
	fmt.Println("max=",max,"min=",min)
}
