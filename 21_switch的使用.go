package main

import "fmt"

func main(){
	var num int
	fmt.Println("请按下楼号：")
	fmt.Scan(&num)
	switch num{//switch 后面写的是变量本身
	case 1:
		fmt.Printf("按下的是%d楼\n",num)
		//break //go语言保留了break关键字，跳出switch语句，不写，默认包含
		fallthrough //不跳出switch语句，后面的无条件执行
	case 2:
		fmt.Printf("按下的是%d楼\n",num)
		//break
		fallthrough
	case 3:
		fmt.Printf("按下的是%d楼\n",num)
		//break
		fallthrough
	case 4:
		fmt.Printf("按下的是%d楼\n",num)
		//break
		fallthrough
	//其他情况
	default:
		fmt.Printf("按下的是%d楼",num)
		//break

	}
}
