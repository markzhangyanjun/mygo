package main

import (
	"fmt"
	"time"
)

func main(){
	i := 0
	for {//for 后面不写任何东西，这个循环条件永远为真，死循环
	   i ++
	   time.Sleep(time.Second)
	   if i == 5{
	   	fmt.Println("i = ",i)
	   	continue //结束本次循环，继续执行下次循环
	   }
	   if i == 6 {
	   	fmt.Println("i = ",i)
	   	break //跳出循环，如果嵌套多个循环那么跳出最近的循环
	   }
	   fmt.Println("i = ",i)

	}
}
