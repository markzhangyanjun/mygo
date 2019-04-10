package main

import "fmt"

func testIb(){

	fmt.Println("aaaaaaaaa")

}
func testIIb(x int){

	//设置recover
	defer func(){
		//recover()
		// 可以打印panic的错误信息
		//fmt.Println(recover())

		if err := recover(); err != nil{
			fmt.Println(err)

		}
	}()


	var a [10]int
	a[x] = 111 //当x为20时候，导致数组越界，产生一个Panic导致程序崩溃

}

func testIIIb(){

	fmt.Println("ccccccccc")

}



func main() {

	testIb()
	testIIb(20)
	testIIIb()


}

