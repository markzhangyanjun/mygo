package main

import "fmt"

func main(){
	//for 初始条件 ； 条件判断 ；条件变化{}
	//1.初始化条件 i ：=1
	//2. 判断条件是否为真，i <=100 ,如果条件为真，执行循环体，如果为假跳出循环
	//3.条件变化
	//4.重复 2，3，4
	sum := 0
	for  i := 1 ; i <= 100; i ++ {
		sum = sum + i
	}
	fmt.Println("sum=",sum)
}