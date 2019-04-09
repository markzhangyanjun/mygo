package main //必须


import (
	io "fmt" //起别名
	"os"
	_ "os" //忽略该包

) //导入包，必须使用，否则编译不过

func main(){
	io.Println("this is a test")
	io.Println("os.ARGS",os.Args)
}




