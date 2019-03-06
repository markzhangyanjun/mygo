package main

import "fmt"

func main(){
	//goto可以用在任何地方但是不能跨函数使用
	fmt.Println("sdfghj")
	goto End
	fmt.Println("aaaaaaa")

End:
	fmt.Println("zyj")

}
