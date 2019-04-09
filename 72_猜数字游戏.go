package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int){

	//设置种子
	rand.Seed(time.Now().UnixNano())

    var num int

	for {
		num = rand.Intn(10000)
		//fmt.Println("num=",num)
		if num >= 1000{
			break
		}
	}

    *p = num

}

func GetNum(keySlice []int,randNum int){

	//取千位
	keySlice[0] = randNum /1000
	//取百位
	keySlice[1] = randNum%1000/100
	//取十位
	keySlice[2] = randNum%1000%100/10
	//取个位
	keySlice[3] = randNum%10

}


func OnGame(randSlice []int){
	var num int
	keySlice := make([]int,4)
	for{
		fmt.Println("请输入一个四位数")
		fmt.Scan(&num)
		//999 < num < 100000
		if 999 <num &&num <= 10000{
			break
		}
		fmt.Println("输入的数不符合要求")
	}
	fmt.Println("num=",num)


	GetNum(keySlice,num)
	//fmt.Println("keySlince=",keySlice)

	n := 0
	for i :=0;i < 4; i++{

		if keySlice[i] > randSlice[i]{
			fmt.Printf("第%d位大了一点\n",i+1)
		}else if keySlice[i]<randSlice[i]{
			fmt.Printf("第%d位小了一点\n",i+1)

		}else if keySlice[i] == randSlice[i]{
			fmt.Printf("第%d位猜对了\n",i+1)
			n ++
		}
		if  n == 4{
			fmt.Println("全猜对了")

	}


	}


}


func main(){

	var randNum int

	//产生一个4位的随机数
	CreatNum(&randNum)

	fmt.Println("randNum=",randNum)

	randSlice :=make([]int,4)

	GetNum(randSlice,randNum)

	fmt.Println("randSlice=",randSlice)


	OnGame(randSlice) //游戏


	//保存这4位数的每一位
	//n := 1234/1000
	//fmt.Println("n=",n)
	//n1 :=1234 %1000
	//fmt.Println("n1=",n1)
	//n2 :=1234 %1000%100
	//fmt.Println("n2=",n2)



}
