package main

import "fmt"

//回调函数，函数有一个参数是函数类型，这就是回调函数
//多态，多种形态，调用同一个接口不同的表现，可以实现不同的表现
//先有想法，后面再实现功能
func add(args ...int)(rest int){
	for _, data := range args{
		rest += data
	}
	return
}
func minus(args ...int)(res int){
	for _, data := range args{
		res = data - res
	}
	return
}
func plus(args ...int)(res int){
	res = 1
	for _ , data := range args{
		res = data * res
	}
	return
}
func div(args ...int)(res int){
	res = 1
	for _ ,data :=range args{
		res = res / data
	}
	return
}
type Functype  func(args ...int)int
//计算机，可以进行四则运算
func Calc(fTest Functype,args ...int)(res int){
	res  = fTest(args...)
	return
}

func main(){

	res := Calc(add,1,2,3,4,5)
	fmt.Println("res = ",res)

	res = Calc(plus,1,2,3)
	fmt.Println("res = ",res)

	res = Calc(div ,10,2,1)
	fmt.Println("res = ",res)
}