package main

import "fmt"

func testI(){

	fmt.Println("aaaaaaaaa")

}
func testII(){

	//fmt.Println("bbbbbbbbb")

	//显示调用panic函数，导致程序终断
	panic("this is panic test")

}
func testIII(){

	fmt.Println("ccccccccc")

}



func main() {

	testI()
	testII()
	testIII()


}
