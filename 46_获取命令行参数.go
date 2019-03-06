package main

import (
	"fmt"
	"os"
)

func main() {
	//接受用户传递的参数，都是以字符串方式传递的
	list :=os.Args
	n := len(list)
	fmt.Println("n=",n)
}
