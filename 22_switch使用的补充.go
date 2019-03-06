package main

import "fmt"

func main(){
	//支持一个初始化语句，初始化语句和变量本身以分号分割
	switch num :=3 ; num{//switch 后面写的是变量本身
	case 1:
		fmt.Printf("按下的是%d楼\n",num)
	case 2:
		fmt.Printf("按下的是%d楼\n",num)
		//break

	case 3,4,5:
		fmt.Printf("按下的是%d楼yyy\n",num)
	case 6:
		fmt.Printf("按下的是%d楼\n",num)
		//break

	//其他情况
	default:
		fmt.Printf("按下的是%d楼",num)

	}

	score :=85
	switch {//可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("优秀")
	case score > 80 :
		fmt.Println("良好")
	case score > 70 :
		fmt.Println("一般")
	case score > 60 :
		fmt.Println("及格")
	}

}