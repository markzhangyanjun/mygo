package main

import "fmt"

func testIa(){

	fmt.Println("aaaaaaaaa")

}
func testIIa(x int){

	var a [10]int
	a[x] = 111 //当x为20时候，导致数组越界，产生一个Panic导致程序崩溃

}

func testIIIa(){

	fmt.Println("ccccccccc")

}



func main() {

	testIa()
	testIIa(20)
	testIIIa()


}
