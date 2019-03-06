package main

import "fmt"

func main(){
	//1
	a := 10
	if a == 10 {
		fmt.Println("a=10")
	}else{//else后面没有条件
		fmt.Print("a !=10")
	}

	//2
	if a :=10;a == 10 {
		fmt.Println("a=10")
	}else { //else后面没有条件
		fmt.Print("a !=10")
	}

	//3
	//这种方法好，如果一个条件判断成功，其他的条件将不会判断，效率高
	a = 10
	if a == 10{
		fmt.Println("a=10")
	}else if a >10{
		fmt.Println("a>10")
	}else if a <10 {
		fmt.Println("a<10")
	}else{
		fmt.Println("这是不可能的")
	}

	//4
	b := 10
	if b == 10{
		fmt.Println("b=10")
	}
	if b > 10{
		fmt.Println("b>10")
	}
	if b <10 {
		fmt.Println("b<10")
	}
}
