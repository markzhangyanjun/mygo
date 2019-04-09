package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	//设置种子，只需要一次
	//如果种子参数一样，每次运行产生的随机数一样
	rand.Seed(time.Now().UnixNano())

	//产生随机数

	for i :=0;i<5;i ++{
		fmt.Println("rand=",rand.Int())//随机很大的数
	}

	for i :=0;i<5;i ++{
		fmt.Println("rand=",rand.Intn(100))//限制在100内的数
	}

}
