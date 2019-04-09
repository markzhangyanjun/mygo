package main

import "fmt"

func main(){
	//有多少个[]就是多少维
	//有多少个[]就用多少个循环

	var a [3][4]int

	k := 0

	for i :=0;i < 3; i++{
		for j :=0;j<4;j++{
			k ++
			a[i][j]= k
			fmt.Printf("a[%d][%d]=%d ",i,j,a[i][j])

		}
		fmt.Println("\n")
	}

	fmt.Println(a)
	//有3个元素，每个元素又是一维数组[4]int
	b := [3][4]int{
		{1,2,3,4},
		{4,5,6,7},
		{8,9,10,11},
	}

	fmt.Println("b=",b)


	//部分初始化,没有初始化的部分的值为0
	c := [3][4]int{
		{1,2,3},
		{4,5,6},
		{8,9},
	}
	fmt.Println("c=",c)


	d := [3][4]int{
		{1,2,3},
		{4,5,6},

	}
	fmt.Println("d=",d)


	//指定只初始化某一个
	e := [3][4]int{
		1:{1,2,3},
	}
	fmt.Println("e=",e)







}