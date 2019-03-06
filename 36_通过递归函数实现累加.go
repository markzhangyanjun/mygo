package main

import "fmt"

func fun()(sum int){
	for i := 1; i <=100; i ++{
		sum += i
	}
	return
}
//sum初始值为0
func fun01(i int)(sum int){
	if  i == 1{
		fmt.Println("i=",i)
		sum = sum + i
		return
	}
	sum = fun01(i-1) + i
	return
}

func fun02()(res bool){
	return
}



func main(){
   var sum  int
   sum  = fun()
   fmt.Println("sum=",sum)

   sum = fun01(100)
	fmt.Println("sum=",sum)

   res := fun02()
   fmt.Println("res=",res)

}
