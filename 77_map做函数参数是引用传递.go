package main

import "fmt"

//map做函数参数是引用传递

func htest(m map[int]string){

	delete(m,1)
}

func main(){

	m := map[int]string{1:"mike",2:"zyj",3:"www"}
	fmt.Println("m=",m)

	htest(m) //在函数内部删除某个key
	fmt.Println("m=",m)

}