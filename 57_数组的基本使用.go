package main

import "fmt"

func main(){
	//定义一个数组 [10]int ,[5]int 是不同类型
	//【数字】，这个数字作为数组元素个数
	var a [10]int
	var b [5]int

	fmt.Printf("len(a)=%d,len(b)=%d\n",len(a),len(b))

	//数组元素个数必须是常量
	//n :=10
	//var c [n]int //non-constant array bound n

	//操作数组元素个数，从0开始，到len()-1，不对称元素，这个数字叫下标
	//这是下标，可以是变量或常量
	a[0]=1
	i :=1
	a[i] =2

	for i :=1;i<len(a);i ++{
		a[i]=i +1

	}

	//打印
	//迭代打印每个元素，默认返回2个值：一个是元素位置，一个是元素本身
	for i,data :=range a{
		fmt.Printf("a[%d]=%d\n",i,data)
	}



}
